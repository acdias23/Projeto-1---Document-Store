# Projeto-1---Document-Store

## Integrantes
Ana Clara Dias Cruz RA: 22.122.073-4
<br>
Carolina Stancov RA: 22.122.017-1
<br>
Gabrielly Hungaro RA: 22.122.059-3
<br>
Sofia Fernandes RA: 22.122.082-5

## Instruções para rodar o banco
```

```

## Descrição das Collections
```
{
    "aluno": {
      "nome": "string",
      "id": "string",
      "curso_id": "string",
      "date": "string"
    },
    "curso": {
        "nome": "string",
        "id": "string",  
        "departamento_id": "string"  
    },
    "departamento": {
        "nome": "string",
        "id": "string",  
        "chefe_id": "string"  
    },
    "disciplina": {
        "nome": "string",
        "id": "string",  
        "curso_id": "string",
        "professor_id": "string" 
    },
    "disciplina_ministrada": {
        "nome": "string",
        "disciplina_id": "string",  
        "professor_id": "string",
        "semestre": "int",
        "ano": "int" 
    },
    "grupo_tcc": {
        "id": "string",
        "aluno1_id": "string",  
        "aluno2_id": "string",
        "aluno3_id": "string",
        "orientador": "string" 
    },
    "historico_escolar": {
        "id": "string",
        "aluno_id": "string",  
        "disciplina_id": "string",
        "semestre": "int",
        "ano": "int",
        "nota_final": "int" 
    },
    "matriz_curricular": {
        "id": "string",
        "curso_id": "string",  
        "disciplina_id": "string",
        "semestre": "int"
    },
    "professor": {
        "id": "string",
        "nome": "string",  
        "departamento_id": "string"
    }
  }

```
