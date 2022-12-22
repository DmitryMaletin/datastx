grammar LookML;

model
  : 'model:' ID '{' view+ '}'
  ;

view
  : 'view:' ID '{' dimension+ measure+ '}'
  ;

dimension
  : 'dimension:' ID '{' field+ '}'
  ;

dimension_group
  : 'dimension_group:' ID '{' field+ '}'
  ;
  
measure
  : 'measure:' ID '{' field+ '}'
  ;

field
  : 'field:' ID '{' field_property+ '}'
  ;

field_property
  : 'type:' field_type
  | 'sql:' STRING
  | 'description:' STRING
  | 'group_label:' STRING
  ;

field_type
  : 'string'
  | 'number'
  | 'boolean'
  | 'date'
  | 'datetime'
  | 'time'
  | 'duration'
  ;

ID  : [a-zA-Z_] [a-zA-Z_0-9]*;
STRING : '"' .*? '"' ;
WS  : [ \t\r\n]+ -> skip;
