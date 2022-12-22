grammar Jinja;

template : block+;

block : for_block | if_block | variable | text;

for_block : '{%' 'for' ID 'in' ID '%}' block+ '{%' 'endfor' '%}';

if_block : '{%' 'if' ID '%}' block+ ('{%' 'elif' ID '%}' block+)* ('{%' 'else' '%}' block+)? '{%' 'endif' '%}';

variable : '{{' ID '}}';

text : ~('{%' | '{{') .*?;

ID : [a-zA-Z][a-zA-Z0-9_]*;
