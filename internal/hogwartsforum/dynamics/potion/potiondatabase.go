package potion

import "html/template"

// declare a map of potion names to potion icons
var PotionIcons = map[string]template.HTML{
	"Agudizadora del Ingenio":             `<img src="https://i.imgur.com/GanNYpr.png" class="tooltip" title="Agudizadora del Ingenio"/>`,
	"Amortentia":                          `<img src="https://i.imgur.com/yuVxqDZ.png" class="tooltip" title="Amortentia"/>`,
	"Antídoto de Glumbumble":              `<img src="https://i.imgur.com/9RBZCeu.png" class="tooltip" title="Antídoto de Glumbumble"/>`,
	"Antídoto para venenos comunes":       `<img src="https://i.imgur.com/0RRcrav.png" class="tooltip" title="Antídoto para Venenos Comunes"/>`,
	"Antídoto para Venenos Tóxicos":       `<img src="https://i.imgur.com/m2sCKZf.png" class="tooltip" title="Antídoto para Venenos Tóxicos"/>`,
	"Brebaje Parlanchín":                  `<img src="https://i.imgur.com/YHbpTZb.png" class="tooltip" title="Brebaje Parlanchín"/>`,
	"Crecepelo":                           `<img src="https://i.imgur.com/lcAX0qO.png" class="tooltip" title="Crecepelo"/>`,
	"Curadora de Furúnculos":              `<img src="https://i.imgur.com/lGHayif.png" class="tooltip" title="Curadora de Furúnculos"/>`,
	"Doxycida":                            `<img src="https://i.imgur.com/ACnnCYS.png" class="tooltip" title="Doxycida"/>`,
	"Elixir para inducir Euforia":         `<img src="https://i.imgur.com/qigGa6E.png" class="tooltip" title="Elixir para inducir Euforia"/>`,
	"Embellecedora":                       `<img src="https://i.imgur.com/CQSs7AL.png" class="tooltip" title="Embellecedora"/>`,
	"Esencia de Díctamo":                  `<img src="https://i.imgur.com/DaWtYLH.png" class="tooltip" title="Esencia de Díctamo"/>`,
	"Esencia de Murtlap":                  `<img src="https://i.imgur.com/TPPuHVj.png" class="tooltip" title="Esencia de Murtlap"/>`,
	"Felix Felicis":                       `<img src="https://i.imgur.com/o54Rv3c.png" class="tooltip" title="Felix Felicis"/>`,
	"Fertilizante de Estiércol de Dragón": `<img src="https://i.imgur.com/z7biKGy.png" class="tooltip" title="Fertilizante de Estiércol de Dragón"/>`,
	"Filtro de Muertos en Vida":           `<img src="https://i.imgur.com/HSFCLlF.png" class="tooltip" title="Filtro de Muertos en Vida"/>`,
	"Filtro de Paz":                       `<img src="https://i.imgur.com/Tm4fAWN.png" class="tooltip" title="Filtro de Paz"/>`,
	"Grasp of Death":                      `<img src="https://i.imgur.com/TA2nuNv.png" class="tooltip" title="Grasp of Death"/>`,
	"Multijugos":                          `<img src="https://i.imgur.com/1kEPEYJ.png" class="tooltip" title="Multijugos"/>`,
	"Pócima para Dormir":                  `<img src="https://i.imgur.com/Z7R6IF4.png" class="tooltip" title="Pócima para Dormir"/>`,
	"Poción Aliento de Fuego":             `<img src="https://i.imgur.com/JiZi0Wh.png" class="tooltip" title="Poción Aliento de Fuego"/>`,
	"Poción Calmante":                     `<img src="https://i.imgur.com/HTMjhyQ.png" class="tooltip" title="Poción Calmante"/>`,
	"Poción de Despertares":               `<img src="https://i.imgur.com/lfgDzmp.png" class="tooltip" title="Poción de Despertares"/>`,
	"Poción de Erumpent":                  `<img src="https://i.imgur.com/TzFgk4h.png" clas="tooltip" title="Poción de Erumpent"/>`,
	"Poción de la Memoria":                `<img src="https://i.imgur.com/ldkNgXk.png" class="tooltip" title="Poción de la Memoria"/>`,
	"Poción de la Risa":                   `<img src="https://i.imgur.com/rnrljbI.png" class="tooltip" title="Poción de la Risa"/>`,
	"Poción del Olvido":                   `<img src="https://i.imgur.com/qsTMBQd.png" class="tooltip" title="Poción del Olvido"/>`,
	"Poción Herbicida":                    `<img src="https://i.imgur.com/JE6nCfo.png" class="tooltip" title="Poción Herbicida"/>`,
	"Poción Herbovitalizante":             `<img src="https://i.imgur.com/FiunEWk.png" class="tooltip" title="Poción Herbovitalizante"/>`,
	"Poción Oculus":                       `<img src="https://i.imgur.com/0IXlBh8.png" class="tooltip" title="Poción Oculus"/>`,
	"Poción para el Dolor de Estómago":    `<img src="https://i.imgur.com/k3cFa0J.png" class="tooltip" title="Poción para el Dolor de Estómago"/>`,
	"Poción Pimentónica":                  `<img src="https://i.imgur.com/U5zJMuB.png" class="tooltip" title="Poción Pimentónica"/>`,
	"Poción Protectora contra las Llamas": `<img src="https://i.imgur.com/dx4Dh3y.png" class="tooltip" title="Poción Protectora contra las Llamas"/>`,
	"Poción Vigorizante":                  `<img src="https://i.imgur.com/V5mkB3V.jpg" class="tooltip" title="Poción Vigorizante"/>`,
	"Poción Volubilis":                    `<img src="https://i.imgur.com/U3SIfC6.png" class="tooltip" title="Poción Volubilis"/>`,
	"Solución Agrandadora":                `<img src="https://i.imgur.com/zWxv0P1.png" class="tooltip" title="Solución Agrandadora"/>`,
	"Solución Encogedora":                 `<img src="https://i.imgur.com/nO9UMLH.png" class="tooltip" title="Solución Encongedora"/>`,
	"Veritaserum":                         `<img src="https://i.imgur.com/29AFHVA.png" class="tooltip" title="Veritaserum"/>`,
	"Zumo de Mandrágora":                  `<img src="https://i.imgur.com/la020XU.jpg" class="tooltip" title="Zumo de Mandrágora"/>`,
}

var PotionIngredients = map[string]string{
	"Veritaserum":                         "Un pelo de cola de Unicornio adulto macho\nPluma de Fénix\nMedio litro de agua del Río Nilo (Egipto)\nUn trozo de dedo de un Grindylow\nCorazón de dragón\nAcónito\nJarabe de eléboro",
	"Felix Felicis":                       "Huevo de Ashwinder\nDrimia maritima\nTentáculo de Murtlap\nTomillo\nCáscara de huevo de Occamy\nRuda\nRábano de caballo",
	"Amortentia":                          "Asfódelo cortado\nTisana\nSemillas de anís verde\nRaíz de Angélica\nComino\nHinojo\nAcónito\nAjenjo",
	"Esencia de Murtlap":                  "Tensa\nEncurtido de tentáculos de murtlap\nMuerdago\nValeriana",
	"Filtro de Muertos en Vida":           "Ajenjo\nAsfódelo\nRaíces de valeriana\nJugo de 12 Granos de sopóforo\nCerebro de perezoso",
	"Multijugos":                          "Sanguijuelas\nCrisopos\nDescurainia sophia\nCentinodia\nPolvo de cuerno de bicornio\nPiel de serpiente arbórea africana\nAlgo de la persona en la que se vaya a convertir",
	"Agudizadora de Ingenio":              "Escarabajos machacados\nBilis de armadillo\nRaíz de jengibre cortada",
	"Antídoto de Glumbumble":              "Melaza de Glumbumble",
	"Embellecedora":                       "Alas de hada\nRocío de la mañana\nPétalos de Rosa perfecta\n4 Medidas de pie de león cortado\nMechón de pelo de unicornio\nRaíz de jengibre",
	"Esencia de Díctamo":                  "Saliva de Salamandra\nDictamo (Planta)",
	"Filtro de Paz":                       "Jarabe de Eléboro\nOpalo/piedra lunar/piedra de deseo\nCuerno de unicornio (polvo)\nPuas de puercoespin (polvo)",
	"Solución Encogedora":                 "Raíces de Margarita\nHigos secos\nOrugas lanudas\nBazo de rata\nSanguijuelas\nCicuta virosa\nAjenjo",
	"Zumo de Mandrágora":                  "Mandragora",
	"Aliento de Fuego":                    "Menta\nValeriana\nSemilla de fuego\nCuerno de dragón en polvo\nLavanda",
	"Antidoto para Venenos Comunes":       "Rocío de luna\nEsporas de vainilla de viento\nMoco de gusarajo\nAcónito\nPiel de serpiente arbórea africana\nAguamiel\nMenta\nMandrágora cocida\nEsencia de lavanda",
	"Antídoto para Venenos Toxicos":       "Semillas de fuego\nAguijones de Billywig\nCuerno de graphorn en polvo\nCaparazones de chizpurfle",
	"Brebaje Parlanchin":                  "Ramitas de valeriana\nAcónito\nDíctamo",
	"Poción Calmante":                     "Lavanda\nCorazón de cocodrilo\nMenta",
	"Poción Crecepelo":                    "Cola de rata\nPúas de puercoespín\nAguijones de billywig",
	"Doxycida":                            "5 Medidas de Bundimun\n1 Hígado de dragón\n3 Caparazones de streeler\n5 Medidas de esencia de cicuta virosa\n3 Medidas de esencia de cicuta\n3 Medidas de tintura de potentilla",
	"Elixir para Inducir Euforia":         "Ramitas de menta\nHigos secos\nPúas de puercoespín\nAjenjo\nSemillas de ricino\nGranos de sopóforo",
	"Fertilizante de Estiercol de Dragón": "Cerebro de perezoso\nCaballitos de mar voladores\nEstiércol de dragón\nMandrágora cocida\nRaíces de margarita\nBazos de rata\nTórax de libélula",
	"Pócima para Dormir":                  "4 ramitas de Lavanda\n6 medidas del ingrediente estándar\n2 cucharadas de moco de gusarajo\n4 ramitas de valeriana",
	"Poción de Despertares":               "6 colmillos de serpiente\n4 medidas de ingrediente estándar\n6 aguijones de billywig secos\n2 ramitas de acónito",
	"Poción de la Memoria":                "1 Pluma de jobberknoll\n3 Galanthus nivalis\n2 Medidas de mandrágora cocida\n2 Medidas de salvia\nHojas de alihotsy\nMenta\nOjos de anguila",
	"Poción de la Risa":                   "Agua de manantial\nHojas de alihotsy\nAlas de billywig\n3 Púas de knarl\nPelo de puffskein\nPolvo de rábano picante\nRisa",
	"Poción del Olvido":                   "2 Gotas de agua del río Lethe\n2 Ramitas de valeriana\n2 Medidas de Ingrediente estándar\n4 Bayas de muérdago",
	"Poción Oculus":                       "Ajenjo\nCuerno de unicornio\nPolvo de ópalo\nMandrágora cocida",
	"Poción Herbicida":                    "Ortigas secas\nPúa de puercoespín\nColmillos de serpiente",
	"Poción Herbovitalizante":             "Corteza de azarollo\nMoly\nDíctamo\nUna pinta de zumo de horklump\n2 Gotas de moco de gusarajo\n7 Colmillos de chizpurfle\nBaba de aguijón de billywig\nUna ramita de menta\nZumo de bayaboom\nMandrágora cocida\nGotas de aguamiel\nMucosa de cerebro de perezoso\nGotas de rocío de luna\nAsfódelo\nCuerno de unicornio\nAcónito\nSangre de salamandra\n10 Espinas de pez león",
	"Poción Protectora Contra las Llamas": "Hongo explosivo\nSangre de salamandra\nPolvos verrugosos",
	"Poción Volubilis":                    "Aguamiel\nRamitas de menta\nMandrágora cocida\nJarabe de eléboro",
	"Solución Agrandadora":                "3 Ojos de pez globo\n1 Bazo de murciélago\n2 Cucharadas de ortigas secas",
	"Poción Curadora de Furúnculos":       "4 babosas cornudas\n2 púas de puercoespín\n6 colmillos de serpiente\nCebollas malolientes\nMoco de gusarajo\nRaíz de jengibre\nEspinas de shrake\nOrtiga seca",
	"Poción Pimentónica":                  "Cuerno de bicornio\nRaíz de mandrágora\nImpatiens capensis",
	"Poción Vigorizante":                  "Hojas de Alihotsy\nAguijones secos de billywig\nMenta\nMandrágora cocida\nInfusión de Ajenjo\nAguamiel\nInfusión de verbena\nCoclearia\nLigústico",
	"Poción de Erumpent":                  "Cuerno de Erumpent\nPolvo de cuerno de unicornio\nEstómago de Gusano Cornudo",
	"Grasp of Death":                      "500gr de semillas de granada\n250 gr de Aceite de díctamo blanco (inflamable)  \n10gr de Polvillo de alas de Hada\n200 gr Jarabe de eléboro\n1 pieza de corteza del árbol vitalizante\n500gr de Rocío de la Mañana\nLágrima de un unicornio",
	"Poción para el dolor de estómago":    "Menta\nJarabe de Menta\nAnís\nPétalos de Manzanilla\nHígado de dragón",
}
