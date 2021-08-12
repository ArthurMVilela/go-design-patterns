# Builder

Builder é um padrão que permite criar objetos passo a passo.

- **Problema:** Classe que precisa ser incializada com 
muitos parametros.
- **Solução:** A construção (incialização) da classe é
removida dela e é movida para outra classes contrutoras
(builders), que implementam etapas para criar uma instância
da classe.

![Exemplo classe builder](https://refactoring.guru/images/patterns/diagrams/builder/solution1.png?id=8ce82137f8935998de80)

