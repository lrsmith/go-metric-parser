# go-metric-parser



## Wavefront WQL
* https://docs.wavefront.com/wavefront_data_format.html#wavefront-data-format-fields : https://docs.wavefront.com/wavefront_data_format.html#wavefront-data-format-fields

# Example Tokenization
```
% go test                                                              
tokens :  [0:TS:ts 2:LPAREN:( 3:IDENT:alpha 9:AND:and 13:NOT:not 17:IDENT:omega 22:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:IDENT:alpha 9:AND:and 13:IDENT:omega 18:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:DQUOTE:" 4:IDENT:~alpha.beta.omega 21:DQUOTE:" 22:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:DQUOTE:" 4:IDENT:alpha.*.omega 17:DQUOTE:" 18:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:DQUOTE:" 4:IDENT:alpha.beta.omega 20:DQUOTE:" 21:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:DQUOTE:" 4:IDENT:alpha 9:DQUOTE:" 10:RPAREN:)]
tokens :  [0:TS:ts 2:LPAREN:( 3:IDENT:alpha 8:RPAREN:)]
```
