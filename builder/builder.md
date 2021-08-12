# Builder

Builder é um padrão que permite criar objetos passo a passo.

- **Problema:** Classe que precisa ser inicialização com 
muitos parametros.
- **Solução:** A construção (incialização) da classe é
removida dela e é movida para outras classes contrutoras
(builders), que implementam etapas para criar uma instância
da classe.

![Exemplo classe com muitos parametros para inicialização](https://refactoring.guru/images/patterns/diagrams/builder/problem2.png?id=2e91039b6c7d2d2df6ee)

![Exemplo classe builder](https://refactoring.guru/images/patterns/diagrams/builder/solution1.png?id=8ce82137f8935998de80)

