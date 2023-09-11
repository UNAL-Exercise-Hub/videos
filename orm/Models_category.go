package orm

import "gorm.io/gorm"

type Musculos struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_musculos;"`
}

type Grupo struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_grupo;"`
}

type Objetivo struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_objetivo;"`
}

type Dificultad struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_dificultad;"`
}

type Equipamento struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_equipamento;"`
}

type Disciplina struct {
	gorm.Model
	Nombre      string
	Descripcion string
	Videos      []*Video `gorm:"many2many:rel_disciplina;"`
}

func default_cat() {
	musculos := []*Musculos{
		{Nombre: "Pectorales", Descripcion: desc_Pectoral},
		{Nombre: "Trapecios", Descripcion: desc_Trapecio},
		{Nombre: "Cuádriceps", Descripcion: desc_Cuadricep},
		{Nombre: "Abdominales", Descripcion: desc_Abdominal},
		{Nombre: "Bíceps", Descripcion: desc_Biceps},
		{Nombre: "Tríceps", Descripcion: desc_Triceps},
		{Nombre: "Deltoides", Descripcion: desc_Deltoides},
	}

	grupos := []*Grupo{
		{Nombre: "Pecho", Descripcion: desc_Pecho},
		{Nombre: "Espalda", Descripcion: desc_Espalda},
		{Nombre: "Brazos", Descripcion: desc_Brazos},
		{Nombre: "Hombros", Descripcion: desc_Hombros},
		{Nombre: "Muslos", Descripcion: desc_Muslos},
		{Nombre: "Pantorrillas", Descripcion: desc_Pantorrillas},
	}

	objetivo := []*Objetivo{
		{Nombre: "Perder o mantener peso", Descripcion: desc_Peso},
		{Nombre: "Salud", Descripcion: desc_Salud},
		{Nombre: "Alivio de estrés", Descripcion: desc_Estres},
		{Nombre: "Fuerza", Descripcion: desc_Fuerza},
	}

	// Descripion taken from https://blog.institutoisaf.es/objetivos-basicos-de-la-rutina-de-entrenamiento-para-principiante
	dificultad := []*Dificultad{
		{Nombre: "Principiante", Descripcion: "Tiempo inferior a 6 meses."},
		{Nombre: "Principiante-Intermedio", Descripcion: "Tiempo superior a 6 meses e inferior a 9 meses."},
		{Nombre: "Intermedio", Descripcion: "Tiempo superior a 9 meses e inferior a 1 año y 6 meses."},
		{Nombre: "Intermedio-Avanzado", Descripcion: "Tiempo superior a 1 año y 6 meses e inferior a 1 año y 9 meses"},
		{Nombre: "Dificil", Descripcion: "Tiempo superior a 1 año y 9 meses"},
	}

	equipamento := []*Equipamento{
		{Nombre: "Colchoneta", Descripcion: "Colchoneta para hacer ejercicios de piso."},
		{Nombre: "Mancuernas", Descripcion: "Mancuernas para añadir peso a los ejercicios."},
		{Nombre: "Maquinas", Descripcion: "Involucra maquinas de ejercicio como trotadoras, bicicletas o maquinas con peso"},
		{Nombre: "Lazo", Descripcion: "Lazo para saltar."},
	}

	disciplina := []*Disciplina{
		{Nombre: "Yoga", Descripcion: desc_Yoga},
		{Nombre: "Crossfit", Descripcion: desc_Crossfit},
		{Nombre: "Zumba", Descripcion: desc_Zumba},
		{Nombre: "Pilates", Descripcion: desc_Pilates},
		{Nombre: "Calistenia", Descripcion: desc_Calistenia},
	}

	db.Create(musculos)
	db.Create(grupos)
	db.Create(dificultad)
	db.Create(objetivo)
	db.Create(equipamento)
	db.Create(disciplina)
}

// Description taken from https://www.fisiocrem.es/blog/conoce-tu-cuerpo/musculos-importantes-cuerpo-humano/
var desc_Pectoral string = "El principal grupo de músculos del pecho son los pectorales. Conseguir un buen desarrollo en estos músculos del cuerpo es muy importante porque fortalece la parte superior interna del cuerpo. El pecho es una parte muy visible y que, además, ayuda a mejorar todos los movimientos y técnicas requeridas en muchas actividades físicas. El press de banca ayuda a trabajarlo."
var desc_Trapecio string = "Ubicado entre los hombros y el cuello se encuentra el trapecio. El trapecio es el encargado de controlar los omóplatos y las escápulas, y es fundamental para movimientos como el encogimiento de hombros y del cuello, lo que nos ayuda a llevar a cabo diferentes movimientos y actividades diarias. También sirve de apoyo para hacer levantamientos por encima de la cabeza. Haciendo elevaciones laterales con mancuernas, los podrás tonificar."
var desc_Cuadricep string = "Son de los músculos del cuerpo más importantes del cuerpo y también los más potentes. Sin ellos no podríamos hacer algo tan básico y necesario como andar. Se encargan de soportar el peso de nuestro cuerpo y ponernos en movimiento. Podemos fortalecerlos con las sentadillas."
var desc_Abdominal string = "Los abdominales son unos de los músculos más trabajados en las salas de gimnasio. Estos ayudan a proteger los órganos de la parte interna y mejoran la respiración. Son un grupo muscular importante para mantener una buena postura y potenciar los movimientos de otros grupos musculares. El mejor ejercicio para ejercitar estos músculos del cuerpo sin lesionar el cuello ni las lumbares es la plancha."
var desc_Biceps string = "Los bíceps se encuentran en la zona superior y delantera del brazo. Estos músculos del cuerpo son los encargados de controlar el movimiento de las articulaciones del codo y del hombro que permite mover los brazos arriba o abajo y de un lado a otro lado. El curl de bíceps con mancuernas es perfecto para trabajarlo."
var desc_Triceps string = "En la parte posterior y superior del brazo se encuentran los tríceps. Estos músculos del cuerpo permiten equilibrar la articulación del hombro y ayudan a encauzar la articulación del codo.El tríceps es muy utilizado para movimientos como lanzar cosas, empujar, escribir y dibujar. Los fondos en barras paralelas o en un escalón harán que los desarrolles."
var desc_Deltoides string = "El deltoides se utiliza en cualquier movimiento que incluya un levantamiento de peso. Proporcionan apoyo al transportar objetos y otras actividades de la vida cotidiana. En los entrenamientos, a la hora de realizar elevaciones laterales y frontales con mancuernas, la función del deltoides es mover el brazo hacia arriba para completar el movimiento."

// Description taken from https://mejorconsalud.as.com/cuales-son-grupos-musculares-como-trabajarlos-rutina-entrenamiento/
var desc_Pecho string = "El músculo principal del grupo muscular del pecho es el pectoral mayor. Una cuestión importante aquí es que sus fibras no están todas alineadas en la misma dirección, por lo que es necesario trabajar de diferentes formas para conseguir un buen desarrollo."
var desc_Espalda string = "El grupo de la espalda está formado por cuatro grandes músculos: trapecio, romboides, dorsal ancho y erector de la columna. Además, hay otros músculos pequeños que también son importantes, como el redondo mayor, el redondo menor  y el infraespinoso."
var desc_Brazos string = "Los brazos están formados por cuatro músculos principales: coracobraquial y braquial anterior (que forman el bíceps braquial), tríceps y antebrazos. En muchas rutinas no se considera el entrenamiento de los antebrazos de forma localizada, sino que se pone el foco en el bíceps o en tríceps."
var desc_Hombros string = "Cuando hablamos de los hombros, nos referimos principalmente al músculo deltoides, que tiene tres cabezas: anterior, media y posterior. Es muy importante desarrollar las tres para no sufrir ningún tipo de desbalance."
var desc_Muslos string = "Los principales músculos de esta sección son los cuádriceps (incluye el vasto lateral, el vasto medial, el vasto intermedio y el recto femoral), los isquiotibiales (semitendinoso, semimembranoso y bíceps femoral) y los glúteos (mayor, menor y medio)."
var desc_Pantorrillas string = "Las pantorrillas están formadas por el gastrocnemio (conocido como «gemelo») y el sóleo. El gastrocnemio es el músculo que se observa claramente en la pantorrilla, mientras que el sóleo es uno profundo, que se halla justo debajo del primero."

// Description taken from https://deportes.upnvirtual.edu.mx/index.php/novedades/blog/468-cuales-son-tus-metas-y-objetivos-al-hacer-ejercicio
var desc_Peso string = "El ejercicio enfocado a la pérdida de peso generalmente es cardiovascular o aeróbico, por ejemplo caminar o trotar de manera moderada o intensa. Son eficientes porque utilizas los músculos más grandes de tu cuerpo aún en periodos cortos. Sin embargo, existen otras formas que también ofrecen resultados para la pérdida de peso y que te permiten modificar la composición corporal."
var desc_Salud string = "El ejercicio enfocado a la pérdida de peso generalmente es cardiovascular o aeróbico, por ejemplo caminar o trotar de manera moderada o intensa. Son eficientes porque utilizas los músculos más grandes de tu cuerpo aún en periodos cortos. Sin embargo, existen otras formas que también ofrecen resultados para la pérdida de peso y que te permiten modificar la composición corporal."
var desc_Estres string = "El ejercicio regular contribuye al bienestar mental, además del físico. Ayuda a tratar la depresión, alivia el estrés y la ansiedad; estimula al cerebro al mejorar la oxigenación del cuerpo. La producción de endorfinas debido al ejercicio reduce los niveles de estrés, muchos se pueden enfocar más o piensan más claramente después de su rutina. También tu confianza en ti mismo puede aumentar, al igual que tu autoestima a medida que mejora tu apariencia por la práctica regular de ejercicio."
var desc_Fuerza string = "El fortalecimiento de huesos y músculos es un objetivo común de la actividad física regular. Las rutinas más comunes se enfocan en ejercicio que se realizan en una posición estática con varias series de repetición como pesas, lo que provoca una mayor resistencia y fortalecimiento a tus músculos. Este tipo de ejercicios previene la osteoporosis porque aumenta y mantiene la masa ósea; el aumento de tu masa muscular mejora tu metabolismo, lo que te ayuda a quemar más calorías."

// Descriptions taken from
// https://www.infobae.com/2013/09/16/1509295-las-cinco-disciplinas-del-fitness-que-son-furor/
var desc_Zumba string = "Fusiona fitness y baile y utiliza un estilo libre coreografiado como método de enseñanza. Representa una actividad aeróbica de dirigidas, utilizando pasos, estilo y música similares a los ritmos latinoamericanos. Sus clases mezclan animados ritmos internacionales con coreografías fáciles de seguir, que hacen un entrenamiento efectivo en todo el cuerpo. "

// https://medlineplus.gov/spanish/ency/patientinstructions/000876.htm#:~:text=El%20yoga%20puede%20mejorar%20el,Ayudarle%20a%20relajarse
var desc_Yoga string = "El yoga es una práctica que conecta el cuerpo, la respiración y la mente. Esta práctica utiliza posturas físicas, ejercicios de respiración y meditación para mejorar la salud general. El yoga se desarrolló como una práctica espiritual hace miles de años. Hoy en día, la mayoría de las personas en occidente que practican yoga lo hacen como ejercicio o para reducir el estrés."

// https://cuidateplus.marca.com/ejercicio-fisico/diccionario/crossfit.html
var desc_Crossfit string = "Sistema de entrenamiento y acondicionamiento en el que se realizan ejercicios funcionales variados a alta intensidad. El objetivo de esta preparación intenso es mejorar las diez capacidades físicas más reconocidas en las pesas: la resistencia cardiovascular, energética, la fuerza, la potencia, la velocidad, la flexibilidad, la precisión, la coordinación, el equilibrio y la agilidad del cuerpo."

// https://efdeportes.com/efd166/ejercitarse-integralmente-el-metodo-pilates.htm
var desc_Pilates string = "Esta técnica fue desarrollada hace más de setenta años por el atleta alemán Joseph Pilates y se trata de un sistema de ejercicios centrado en mejorar la flexibilidad y fuerza para todo el cuerpo sin incrementar su volumen. Más que un entrenamiento físico, el método Pilates utiliza una serie de movimientos controlados atractivos tanto para la mente como para el cuerpo. La técnica Pilates integra teorías occidentales y orientales y relaciona la práctica de ejercicios específicos acoplados con técnicas de respiración."

// https://www.expansion.com/fueradeserie/cuerpo/2022/09/18/630c88f5468aeb68158b462a.html
var desc_Calistenia string = "La calistenia es un sistema de entrenamiento que se especializa en hacer uso del propio peso corporal para realizar los diferentes ejercicios. Muy completo y versátil, este método de ejercitar el cuerpo se ha convertido en una forma de hacer deporte apreciada gracias a su accesibilidad. "
