## Solutions for Wildberries practical tasks 
### Day 1  
#### Задание  
> Реализовать паттерн Фасад в соответствии с конвенцией.  
> Реализация паттерна должна быть следующая:  
> описать пакет facade в директории pkg/facade в директории cmd/facade создать файл main.go,  
> который фактически будет представлять собой интеграционное тестирование пакета.  

#### Решение  
В моём примере реализован интерфейс Planner, через который можно составить план на испытательный срок.


### Day 2
#### Задание
> Реализовать паттерн строитель в соответствии с конвенцией

#### Решение
В моём примере реализован строитель простых сайтов.  
> Для реализации потребуется:
> 1. Пользовательский тип Director, который будет распоряжаться строителем и отдавать ему комманды в нужном порядке, а строитель будет их выполнять.(Director)
> 2. Базовый интерфейс Builder, который описывает команды строителя, которые он обязан выполнять.(siteBuilder)
> 3. Пользовательский тип ConcreteBuilder, который реализует интерфейс строителя и взаимодействует со сложным объектом.(BusinessCardSite/PhotogallerySite)
> 4. Пользовательский тип сложного объекта Product.(Site)
