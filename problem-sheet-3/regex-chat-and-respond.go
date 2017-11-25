package main
// https://code.visualstudio.com/docs/editor/integrated-terminal  in go 
// Regex in go https://github.com/google/re2/wiki/Syntax
// Case insensitive regex in go https://stackoverflow.com/questions/15326421/how-do-i-do-a-case-insensitive-regular-expression-in-go
// https://www.regular-expressions.info/refwordboundaries.html

import (
	"fmt"
	"math/rand"
	"time"
	"regexp" 
	"strings"
)


func ElizaResponse(input string) string{
	
	 // Try changing this number!
	if matched, _ := regexp.MatchString(`.*\bfather\b.*`, input) ; matched {
		return  "Why don't you tell me more about your father?"
	}

	re := regexp.MustCompile(`(.*)\b[Ff]riend\b`)
	if matched := re.MatchString(input) ; matched {
		captured := Reflect(input)
		answer := "Why do you think $1?"
		return fmt.Sprintf(answer, captured)	
	}

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",		
	}

	myAnswers := []string{
		"How do you know you can't $1?",
		"Perhaps you could $1 if you tried.",
		"Did you really try to $1?",
		"What would it take for you to $1?",
		"Is it too difficult for you to $1?",
		"Are you unable to $1 for serious reasons?",
		"Is your lack of posibilities preventing you from $1?",
	}

	myResp := regexp.MustCompile(`[Ii] can't (.*)`)
	if matched := myResp.MatchString(input) ; matched {
		captured := Reflect(input);
		randAns := myAnswers[rand.Intn(len(myAnswers))]			
		return fmt.Sprintf(randAns, captured)
	}


	return answers[rand.Intn(len(answers))]	;
	
}

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println();
	
	fmt.Println("People say I look like both my mother and father.");
	fmt.Println(ElizaResponse("Eliza: People say I look like both my mother and father."));
	fmt.Println();

	fmt.Println("Father was a teacher.");
	fmt.Println(ElizaResponse("Father was a teacher."));
	fmt.Println();

	fmt.Println("I was my father’s favourite.");
	fmt.Println(ElizaResponse("I was my father’s favourite."));
	fmt.Println();

	fmt.Println("I’m looking forward to the weekend.");
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."));
	fmt.Println();

	fmt.Println("My grandfather was French!");
	fmt.Println(ElizaResponse("My grandfather was French!"));
	fmt.Println();
	
	fmt.Println("I am happy.");
	fmt.Println(ElizaResponse("I am happy."));
	fmt.Println();

	fmt.Println("I am not happy with your responses.");
	fmt.Println(ElizaResponse("I am not happy with your responses."));
	fmt.Println();

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.");
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."));
	fmt.Println();


	fmt.Println("I am supposed to just take what you’re saying at face value?");
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"));
	fmt.Println();

	fmt.Println("You are my friend.");
	fmt.Println(ElizaResponse("You are my friend."));
	fmt.Println();
	
	fmt.Println("I can't do this anymore");
	fmt.Println(ElizaResponse("I can't do this anymore"));
	fmt.Println();

	fmt.Println("I can't believe she did this!");
	fmt.Println(ElizaResponse("I can't believe she did this!"));
	fmt.Println();

	fmt.Println("I can't feel any pain right now.");
	fmt.Println(ElizaResponse("I can't feel any pain right now."));
	fmt.Println();
	
}

func Reflect(match string) string{

    // Split is used to break up strings into a list of substrings
	boundaries := regexp.MustCompile(`\n`)
	words := boundaries.Split(match, -1)

    for i, word := range words {
        if reflectedWord, ok := reflections[word]; ok {
            words[i] = reflectedWord
        }
    }

    // Join takes a list of strings and joins them together again
    return strings.Join(words, " ") 
}

	
var reflections = map[string]string{    
	// Synonyms
	"\bi\b": "I",
	"[Ii]'m": "[Ii] am",
	"you're": "you are",
	"[Ii]'ll": "[Ii] will",    
	"you'll": "you will",    
	"[Ii]'d": "[Ii] would",    
	"you'd": "you would",    
	"[Ii]'ve": "[Ii] have",
	"you've": "you have",    
	"[Ii]'ve had": "[Ii] have had",
	"cant": "can't",
	"dont": "don't",
	"wont": "won't",
	"couldnt": "couldn't",
	"didnt": "didn't",
	"wouldnt": "wouldn't",

	// Subject Pronouns
	"[Ii]": "you",
	"he": "she",
	"she": "he",

	// Object Pronouns
	"you": "me",
	"me": "you",
	"him": "her",
	"her": "him",

	// Possesive Adjectives
	"your": "my",
	"my": "your",

	// Possesive Pronouns
	"yours": "mine",
	"mine": "yours",
	"his": "hers",
	"hers": "his",

	// Reflexive Pronouns
	"yourself": "myself",
	"myself": "yourself",
	"himself": "herself",
	"herself": "himself",

	// To be
		// Present
	"am": "are", 
	"are": "am", 
		// Past
	"was": "were",    

	// To Will
		// Present    
	"[Ii] will": "you will",
	"you will": "I will",
		// Past
	"[Ii] would": "you would",
	"you' would": "I would",


	// To Have
		// Present
	"[Ii] have": "you have",
	"you have": "I have",
		// Past
	"[Ii] have had": "you've had", 	
}
