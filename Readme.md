# Probuff

## comando simples: protoc --go_out=. person.proto
### esse comando gera o arqvuio person.pb.go no diretorio que está no person.proto ( mudar ele para uma pasta especifica proto por exemplo ?)


estrutura básica arquivo person.proto:

### Versão
```
syntax = "proto3"; 
```

### Opção de pasta de estino
```
option go_package = "/person"; 
```

### Struct que irá ser criada
```
message Person {
  string name = 1;
  int32 age = 2;
  repeated string emails = 3;
}
```
