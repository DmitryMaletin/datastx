grammar Jinja;

template
  : block+
  ;

block
  : variable
  | for_loop
  | if_statement
  | comment
  ;

variable
  : '{{' variable_expression '}}'
  ;

variable_expression
  : ID ('.' ID)*
  | function_call
  | filter
  ;

function_call
  : ID '(' (variable_expression (',' variable_expression)*)? ')'
  ;

filter
  : variable_expression '|' ID (':' variable_expression)?
  ;

for_loop
  : '{% for' ID 'in' ID '%}' block+ '{% endfor %}'
  ;

if_statement
  : '{% if' variable_expression '%}' block+ ('{% elif' variable_expression '%}' block+)* ('{% else %}' block+)? '{% endif %}'
  ;

comment
  : '{#' .*? '#}'
  ;

ID  : [a-zA-Z_] [a-zA-Z_0-9]*;
WS  : [ \t\r\n]+ -> skip;
