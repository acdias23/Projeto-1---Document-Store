package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/jaswdr/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Aluno struct {
	Name string
	Date string
}

type Professor struct {
	Professor string `bson:"nome"`
}

type Curso struct {
	Nome string `bson:"nome"`
}

type Departamento struct {
	Nome string `bson:"nome"`
}

type Disciplina struct {
	Nome string `bson:"nome"`
}

type DisciplinaMinistrada struct {
	ProfessorID  primitive.ObjectID `bson:"professor_id"`
	DisciplinaID primitive.ObjectID `bson:"disciplina_id"`
	Semestre     int                `bson:"semestre"`
	Ano          int                `bson:"ano"`
}

var cursos = []string{
	"Administracao",
	"Ciencia da Computacao",
	"Engenharia Civil",
	"Engenharia Eletrica",
}

var cursos_id = []primitive.ObjectID{}
var departamentos_id = []primitive.ObjectID{}
var disciplinas_adm_id = []primitive.ObjectID{}
var disciplinas_cc_id = []primitive.ObjectID{}
var disciplinas_engc_id = []primitive.ObjectID{}
var disciplinas_enge_id = []primitive.ObjectID{}
var alunos_id = []primitive.ObjectID{}
var professores_id = []primitive.ObjectID{}

//type alunoData struct {
//	CursoID primitive.ObjectID `bson:"curso_id"`
//CursoID int
//}

func pickRandom(arr []primitive.ObjectID) primitive.ObjectID {
	idx := rand.IntN(len(arr))
	return arr[idx]
}

func randYear() int {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int64N(delta) + min
	return time.Unix(sec, 0).Year()
}

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("").SetServerAPIOptions(serverAPI) //Inserir string de conexão

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("universidade")

	m := make(map[string][]string)

	m["Administracao"] = []string{
		"Matematica aplicada a administracao",
		"Fundamentos da administracao",
		"Estudos em macroeconomia",
		"Sociologia I",
		"Sociologia II",
		"Linguagens e generos textuais",
		"Etica nas organizacoes",
		"Ensino social cristao",
		"Calculo basico I",
		"Calculo basico II",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia I",
		"Filosofia II",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Gramatica",
		"Estudos em microeconomia",
		"Sustentabilidade",
		"Probabilidade estatistica",
		"Estudos do MEI",
		"Gestao de negocios",
		"Gestao de pessoas",
		"Gesta Financeira",
		"Matematica Financeira",
		"Gestao de recusos humanos",
		"PowerBI",
		"Excel Basico",
		"Excel intermediario",
		"Excel avancado",
		"Redacao",
		"Comunicacao para empresas",
		"Estrategia de marketing",
		"Direito Tributario",
		"Logistica",
		"Implementacao de negocios",
	}

	m["Ciencia da Computacao"] = []string{
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Sociologia",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Circuitos eletricos",
		"Ecologia e sustentabilidade",
		"Sistemas embarcados",
		"Redes de computadores",
		"Calculo numerico",
		"Termodinamica",
		"Cinematica e dinamica",
		"Algebra Linear",
		"Automatos",
		"Redes moveis",
		"Algoritmos",
		"Sistema Operacional",
		"Computacao grafica",
		"Compiladores",
		"Empreendedorismo",
		"Inovacao",
		"Teste de software",
		"IA",
		"Orientacao a objetos",
		"Desenvolvimento de jogos",
		"Robotica",
		"TCC I",
		"TCC II",
	}

	m["Engenharia Civil"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Arquitetura e representacao grafica",
		"Mecanica geral",
		"Topografia",
		"Eletricidade geral",
		"Instalacoes eletricas",
		"Economia",
		"Geotecnia I",
		"Geotecnia II",
		"Transportes I",
		"Transportes II",
		"Optativa I",
		"Optativa II",
		"Custos",
		"Tecnologia das construcoes",
		"Metodos estatisticos",
		"Termodinamica",
		"Gestao de obras",
		"Planejamento de obras",
	}

	m["Engenharia Eletrica"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"sinais e Sistemas",
		"Praticas de inovacao I",
		"Praticas de inovacao II",
		"Praticas de inovacao III",
		"Circuitos integrados",
		"Mecanica Geral",
		"Gestao organizacional",
		"TCC 1",
		"TCC 2",
		"Redes",
		"Mecanica dos solidos",
		"Mecanica dos fluidos",
		"Sistemas eletricos",
		"Seguranca do trabalho",
		"Economia",
		"Conversao de energia I",
		"Conversao de energia II",
		"Conversao de energia III",
	}

	departamentos := [5]string{
		"Matematica",
		"Ciencia da Computacao",
		"Quimica",
		"Fisica",
		"Engenharia",
	}

	fakeName := faker.New().Person()
	fakeDate := faker.New().Time()

	var pk primitive.ObjectID
	for i := 0; i < 20; i++ {
		pk = insertAluno(db, Aluno{fakeName.Name(), fakeDate.UnixDate(time.Now())})
		alunos_id = append(alunos_id, pk)
	}

	for i := 0; i < 15; i++ {
		pk = insertProfessor(db, fakeName.Name())
		professores_id = append(professores_id, pk)
	}

	for _, curso := range cursos {
		pk = insertCurso(db, curso)
		cursos_id = append(cursos_id, pk)
	}

	for _, departamento := range departamentos {
		pk = insertDepartamento(db, departamento)
		departamentos_id = append(departamentos_id, pk)
	}
	for _, disc := range m["Administracao"] {
		pk = insertDisciplina(db, disc)
		disciplinas_adm_id = append(disciplinas_adm_id, pk)
	}

	for _, disc := range m["Ciencia da Computacao"] {
		pk = insertDisciplina(db, disc)
		disciplinas_cc_id = append(disciplinas_cc_id, pk)
	}

	for _, disc := range m["Engenharia Civil"] {
		pk = insertDisciplina(db, disc)
		disciplinas_engc_id = append(disciplinas_engc_id, pk)
	}

	for _, disc := range m["Engenharia Eletrica"] {
		pk = insertDisciplina(db, disc)
		disciplinas_enge_id = append(disciplinas_enge_id, pk)
	}

	atualizarFKAluno(db)
	atualizarFKCurso(db)
	atualizarFKDepartamento(db)
	atualizarFKDisciplina(db, disciplinas_adm_id[:], 0)
	atualizarFKDisciplina(db, disciplinas_cc_id[:], 1)
	atualizarFKDisciplina(db, disciplinas_engc_id[:], 2)
	atualizarFKDisciplina(db, disciplinas_enge_id[:], 3)
	atualizarFKProfessor(db)

	inserirDisciplinaMinistrada(db)
	inserirGrupoTCCHistoricoMatriz(db)
}

func inserirDisciplinaMinistrada(db *mongo.Database) {
	collection := db.Collection("disciplina_ministrada")

	inserirDisciplinas := func(disciplinas []primitive.ObjectID) {
		var disciplinasMinistradas []interface{}

		for i, disc := range disciplinas {
			disciplinaMinistrada := bson.D{
				{Key: "professor_id", Value: pickRandom(professores_id)},
				{Key: "disciplina_id", Value: disc},
				{Key: "semestre", Value: (i / 5) + 1},
				{Key: "ano", Value: randYear()},
			}

			disciplinasMinistradas = append(disciplinasMinistradas, disciplinaMinistrada)
		}

		if len(disciplinasMinistradas) > 0 {
			collection.InsertMany(context.Background(), disciplinasMinistradas)
		}
	}

	inserirDisciplinas(disciplinas_adm_id)
	inserirDisciplinas(disciplinas_cc_id)
	inserirDisciplinas(disciplinas_engc_id)
	inserirDisciplinas(disciplinas_enge_id)
}

func inserirGrupoTCCHistoricoMatriz(db *mongo.Database) {

	grupoCollection := db.Collection("grupo_tcc")
	historicoCollection := db.Collection("historico_escolar")
	matrizCollection := db.Collection("matriz_curricular")

	for i := 0; i < 6; i++ {
		grupoTCC := bson.D{
			{Key: "aluno1_id", Value: pickRandom(alunos_id)},
			{Key: "aluno2_id", Value: pickRandom(alunos_id)},
			{Key: "aluno3_id", Value: pickRandom(alunos_id)},
			{Key: "orientador", Value: pickRandom(professores_id)},
		}

		grupoCollection.InsertOne(context.Background(), grupoTCC)
	}

	for _, aluno := range alunos_id {
		var alunoData struct {
			DataEntrada string `bson:"date"`
			CursoID     int    `bson:"curso_id"`
		}

		alunoCollection := db.Collection("aluno")
		alunoCollection.FindOne(context.Background(), bson.M{"_id": aluno}).Decode(&alunoData)

		dateStr := alunoData.DataEntrada
		const layout = "Mon Jan 2 15:04:05 -07 2006"

		t, err := time.Parse(layout, dateStr)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		ano := int64(t.Year())

		var disciplinas []primitive.ObjectID
		switch alunoData.CursoID {
		case 0:
			disciplinas = disciplinas_adm_id
		case 1:
			disciplinas = disciplinas_cc_id
		case 2:
			disciplinas = disciplinas_engc_id
		case 3:
			disciplinas = disciplinas_enge_id
		}

		for i, disc := range disciplinas {
			historico := bson.D{
				{Key: "aluno_id", Value: aluno},
				{Key: "disciplina_id", Value: disc},
				{Key: "semestre", Value: (i / 5) + 1},
				{Key: "ano", Value: ano + int64(i/10)},
				{Key: "nota_final", Value: rand.Float32() * 10},
			}

			historicoCollection.InsertOne(context.Background(), historico)
		}
	}

	for _, curso := range cursos_id {
		var disciplinas []primitive.ObjectID
		switch curso {
		case cursos_id[0]:
			disciplinas = disciplinas_adm_id
		case cursos_id[1]:
			disciplinas = disciplinas_cc_id
		case cursos_id[2]:
			disciplinas = disciplinas_engc_id
		case cursos_id[3]:
			disciplinas = disciplinas_enge_id
		}

		for i, disc := range disciplinas {
			matrizCurricular := bson.D{
				{Key: "curso_id", Value: curso},
				{Key: "disciplina_id", Value: disc},
				{Key: "semestre", Value: (i / 5) + 1},
			}

			matrizCollection.InsertOne(context.Background(), matrizCurricular)
		}
	}
}

func insertAluno(db *mongo.Database, aluno Aluno) primitive.ObjectID {
	collection := db.Collection("aluno")

	result, err := collection.InsertOne(context.Background(), aluno)
	if err != nil {
		return primitive.NilObjectID
	}

	alunoID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID
	}
	return alunoID
}

func atualizarFKAluno(db *mongo.Database) {
	collection := db.Collection("aluno")

	for _, alunoID := range alunos_id {
		cursoID := pickRandom(cursos_id)

		_, _ = collection.UpdateOne(
			context.Background(),
			bson.M{"_id": alunoID},
			bson.M{
				"$set": bson.M{
					"curso_id": cursoID,
				},
			},
		)

	}
}

func insertProfessor(db *mongo.Database, nome string) primitive.ObjectID {
	collection := db.Collection("professor")

	professor := bson.D{
		{Key: "nome", Value: nome},
	}

	result, err := collection.InsertOne(context.Background(), professor)
	if err != nil {
		return primitive.NilObjectID
	}

	professorID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID
	}

	return professorID
}

func atualizarFKProfessor(db *mongo.Database) {
	collection := db.Collection("professor")

	for _, professorID := range professores_id {
		departamentoID := pickRandom(departamentos_id)

		_, _ = collection.UpdateOne(
			context.Background(),
			bson.M{"_id": professorID},
			bson.M{
				"$set": bson.M{
					"departamento_id": departamentoID,
				},
			},
		)

	}
}

func insertCurso(db *mongo.Database, nome string) primitive.ObjectID {
	collection := db.Collection("curso")

	curso := bson.D{
		{Key: "nome", Value: nome},
	}

	result, err := collection.InsertOne(context.Background(), curso)
	if err != nil {
		return primitive.NilObjectID
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID
	}

	return insertedID
}

func atualizarFKCurso(db *mongo.Database) {
	collection := db.Collection("curso")

	for _, cursoID := range cursos_id {
		departamentoID := pickRandom(departamentos_id)

		_, _ = collection.UpdateOne(
			context.Background(),
			bson.M{"_id": cursoID},
			bson.M{
				"$set": bson.M{
					"departamento_id": departamentoID,
				},
			},
		)

	}
}

func insertDepartamento(db *mongo.Database, nome string) primitive.ObjectID {
	collection := db.Collection("departamento")

	departamento := bson.D{
		{Key: "nome", Value: nome},
	}

	result, err := collection.InsertOne(context.Background(), departamento)
	if err != nil {
		return primitive.NilObjectID
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID
	}

	return insertedID
}

func atualizarFKDepartamento(db *mongo.Database) {
	collection := db.Collection("departamento")

	for _, departamentoID := range departamentos_id {
		chefeID := pickRandom(professores_id)

		_, _ = collection.UpdateOne(
			context.Background(),
			bson.M{"_id": departamentoID},
			bson.M{
				"$set": bson.M{
					"chefe_id": chefeID,
				},
			},
		)

	}
}

func insertDisciplina(db *mongo.Database, nome string) primitive.ObjectID {
	collection := db.Collection("disciplina")

	disciplina := bson.D{
		{Key: "nome", Value: nome},
	}

	result, err := collection.InsertOne(context.Background(), disciplina)
	if err != nil {
		return primitive.NilObjectID
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID
	}

	return insertedID
}

func atualizarFKDisciplina(db *mongo.Database, arr []primitive.ObjectID, idx int) {
	collection := db.Collection("disciplina")

	for _, disciplinaID := range arr {
		cursoID := cursos_id[idx]
		professorID := pickRandom(professores_id)

		_, _ = collection.UpdateOne(
			context.Background(),
			bson.M{"_id": disciplinaID},
			bson.M{
				"$set": bson.M{
					"curso_id":     cursoID,
					"professor_id": professorID,
				},
			},
		)

	}
}
