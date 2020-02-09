package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var basicRe = regexp.MustCompile(`<div class="des f-cl" [^>]+>(.+) \| ([\d]+)岁 \| (.+) \| (.+) \| ([\d]+)cm \| ([^<]+)</div>`)
var genderRe = regexp.MustCompile(`"genderString":"(.)."`)
var weightRe = regexp.MustCompile(`([\d]+)kg`)
var xinzuoRe = regexp.MustCompile(`"(..座)\(`)
var occupationRe = regexp.MustCompile(`"月收入:[^"]*","([^"]*)"`)
var hukouRe = regexp.MustCompile(`"籍贯:([^"]*)"`)
var carRe = regexp.MustCompile(`"(..车)"`)
var houseRe = regexp.MustCompile(`"(已购房)"`)
var idUrlRe = regexp.MustCompile(`http://album\.zhenai\.com/u/([\d]+)`)

//var guessRe = regexp.MustCompile(`http://album\.zhenai\.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)
	profile.Weight, _ = strconv.Atoi(extractString(contents, weightRe))
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Car = extractString(contents, carRe)
	profile.House = extractString(contents, houseRe)

	extractBasic(contents, basicRe, &profile)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	//matches := guessRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        string(m[1]),
	//		ParserFunc: ProfileParser(string(m[2])),
	//	})
	//}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractBasic(contents []byte, re *regexp.Regexp, profile *model.Profile) {
	match := re.FindSubmatch(contents)

	if len(match) == 7 {
		age, _ := strconv.Atoi(string(match[2]))
		height, _ := strconv.Atoi(string(match[5]))
		profile.Age = age
		profile.Height = height
		profile.Income = string(match[6])
		profile.Marriage = string(match[4])
		profile.Education = string(match[3])
		profile.City = string(match[1])

	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
