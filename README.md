<h3 align="center">Frequency Analysis</h3>

In cryptanalysis, frequency analysis (also known as counting letters) is the study of the frequency of letters or groups of letters in a ciphertext. The method is used as an aid to breaking classical ciphers.

Frequency analysis is based on the fact that, in any given stretch of written language, certain letters and combinations of letters occur with varying frequencies. Moreover, there is a characteristic distribution of letters that is roughly the same for almost all samples of that language. For instance, given a section of English language, E, T, A and O are the most common, while Z, Q, X and J are rare. Likewise, TH, ER, ON, and AN are the most common pairs of letters (termed bigrams or digraphs), and SS, EE, TT, and FF are the most common repeats. The nonsense phrase "ETAOIN SHRDLU" represents the 12 most frequent letters in typical English language text. 

This program will create an HTML document containing a graph with the occurence of each letter in a certain phrase.
![image](https://j.gifs.com/pZPBk1.gif)

For data visualization https://github.com/go-echarts/go-echarts is used.


#### Installation:

`git clone https://github.com/Nisophix/FrequencyAnalysis`

`cd FrequencyAnalysis`

`go mod tidy`

`go run FrequencyAnalysis.go`
