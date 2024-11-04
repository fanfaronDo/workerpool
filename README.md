# WorkerPool

1. Структура Worker

    Структура, которая обрабатывает задачи.
    
    Содержит поля:
    - содржит поле Id, является идентификатором корутины.
    
    Методы: 
    - Proccess принемает данные (строку) иметирует операцию и выводит на эран id, данные

2. Структура WorkerPool

    Управляет добавлением, удалением и выполнением задач.
    
    Содержит поля:
    - workers []*Worker: Слайс, хранящий доступных для обработки задач.
    - tasks chan string: Канал для получения задач, которые необходимо обработать.
    - addWorkerCh chan *Worker: Канал для добавления новых worker в пул.
    - removeWorkerCh chan *Worker: Канал для удаления worker из пула.
    - wg sync.WaitGroup: WaitGroup, для синхронизаци горутин.

    Методы public: 
    - func NewWorkerPool() *WorkerPool - конструктор, для инициализации
    - func (p *WorkerPool) Start()
    - Запускает private метод run, который будет слушать каналы через мультипликсатор select и обрабатывать задачи. 
    - func (p *WorkerPool) AddWorker(worker *Worker). Добавить новый worker в пул 
    - func (p *WorkerPool) RemoveWorker(worker *Worker) удалить worker 
    - func (p *WorkerPool) Submit(task string) отправка задачи в пул. 
    - func (p *WorkerPool) Wait() Метод ожидания завершения всех задач. 


Запуск тестового варианта:
    
    make run

Сборка:

    make build

Запуск тестов:

    make test
