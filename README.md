### Frequency Analysis

In cryptanalysis, frequency analysis (also known as counting letters) is the study of the frequency of letters or groups of letters in a ciphertext. The method is used as an aid to breaking classical ciphers.

Frequency analysis is based on the fact that, in any given stretch of written language, certain letters and combinations of letters occur with varying frequencies. Moreover, there is a characteristic distribution of letters that is roughly the same for almost all samples of that language. For instance, given a section of English language, E, T, A and O are the most common, while Z, Q, X and J are rare. Likewise, TH, ER, ON, and AN are the most common pairs of letters (termed bigrams or digraphs), and SS, EE, TT, and FF are the most common repeats.[1] The nonsense phrase "ETAOIN SHRDLU" represents the 12 most frequent letters in typical English language text. 

This program will create an HTML document containing a graph with the occurence of each letter in a certain phrase.
![image](https://user-images.githubusercontent.com/89792349/154843261-d536124a-95ee-456c-baf0-8c67faf5c758.png)

For data visualization https://github.com/go-echarts/go-echarts is used.


installation:
> `https://github.com/Nisophix/frequencyAnalysis`
>  `cd frequencyAnalysis`
>  `go mod tidy`
>  `go run FrequencyAnalysis.go`
