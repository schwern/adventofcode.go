package day05

import(
    "regexp"
    "github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func hasPair( str string ) bool {
    var prev rune
    for _, letter := range str {
        if letter == prev {
            return true
        }
        prev = letter
    }
    
    return false
}

func IsNice( info string ) bool {    
    // Contains at least 3 vowels
    three_vowels_re := *regexp.MustCompile(`[aeiou]`)
    vowels := three_vowels_re.FindAllString( info, 3 )
    if len(vowels) < 3 {
        return false
    }
    
    // At least one letter that appears twice in a row
    if !hasPair( info ) {
        return false
    }
    
    // It does not contain the strings ab, cd, pq, or xy
    pairs_re := *regexp.MustCompile(`(?:ab|cd|pq|xy)`)
    if pairs_re.MatchString( info ) {
        return false
    }
    
    return true
}

func IsNice2( info string ) bool {
    rules := []pcre.Regexp{ 
        pcre.MustCompile(`(..).*\1`, 0),
        pcre.MustCompile(`(.).\1`, 0),
    }

    for _, rule := range rules {
        matcher := rule.MatcherString( info, 0 )
        if !matcher.Matches() {
            return false
        }
    }
    
    return true
}
