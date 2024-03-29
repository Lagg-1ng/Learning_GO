package main

/*func theory() {
    go sleepyGopher() // Начало горутины
    time.Sleep(4 * time.Second) // Ожидание
} // Здесь все горутины останавливаются
/*
Горутины позволяют сделать любую функцию асинхронной. Другими словами, какую бы функцию ни написал разработчик,
её можно запустить в фоновом режиме, и она будет работать. В то же время планировщик Go сам распределит нагрузку по ядрам, чтобы каждое из них было эффективно нагружено.
Запустить горутину так же просто, как вызвать функцию: нужно поставить ключевое слово go перед вызовом.
*/

//Каналы позволяют передавать и синхронизировать данные между горутинами так, чтобы одни и те же байты не попадали в две разные горутины.
//Для создания канала используется встроенная функция make:

// ~ c := make(chan int)

//Многопоточность

/*
Что такое «поток» в контексте операционной системы?
Поток выполнения (native/kernel thread) — это часть процесса, в которой инструкции могут выполняться независимо и иметь доступ к общим ресурсам.
За управление потоками отвечает планировщик ОС.
Многопоточность — свойство железа и софта, при котором несколько потоков могут выполняться параллельно, не мешая друг другу.
*/

//нестандартная работа с датой

/*input := "2017-08-31"
layout := "2006-01-02"
t, _ := time.Parse(layout, input)
fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017

//В отличие от C, все переменные в Go сразу имеют начальные значения: например, переменные числовых типов равны 0, а строковые переменные содержат пустую строку.

//ООП

/*
Абстракция — возможность определить характеристики (свойства и методы) объекта, которые полностью описывают его возможности.
В Go нет классов, но структуры с методами служат им неплохой заменой.

Инкапсуляция — возможность скрыть реализацию объекта, предоставив пользователю некую спецификацию (интерфейс) взаимодействия с ним.
Go даёт возможность задать область видимости (публичные/приватные) методов структур и позволяет спрятать реализацию.

Наследование — возможность создания производных от родительского объекта, которые будут расширять или изменять свойства и поведение родителя.
К сожалению, Go не реализует в полной мере механизм наследования, но есть встраивание — можно создавать типы на основе существующих.

Полиморфизм — возможность одному и тому же фрагменту кода работать с разными типами данных. Это происходит, когда объект может вести себя как другой объект.
В Go нет полиморфизма в классическом понимании, однако похожие действия можно реализовать с помощью интерфейсов. Интерфейс определяет список методов,
которые должен реализовывать тип, чтобы удовлетворять данному интерфейсу. Это ослабляет строгую типизацию и позволяет передавать в параметрах разные типы данных,
поддерживающие один и тот же интерфейс.
*/

/*
   Рассмотрим язык Go с точки зрения функционального программирования.
   Функции высшего порядка — функции, которые могут в аргументах принимать другие функции и возвращать функции в качестве результата.
   В Go функции рассматриваются как значения и могут передаваться в другие функции, возвращаясь в виде результата.
   Замыкания. Go позволяет определять и использовать функции, которые ссылаются на переменные своей родительской функции.
   Чистые функции. В Go можно определять функции, которые зависят только от входящих аргументов и не влияют на глобальное состояние.
   Рекурсия. Как и в большинстве языков, в Go можно применять рекурсивные вызовы функций.
   Ленивые вычисления. В Go нет поддержки ленивых (отложенных) вычислений.
   Иммутабельность переменных. В Go переменные могут изменять своё значение, поэтому иммутабельность (неизменяемость) переменных отсутствует.
*/

/*
defer — это ещё одна необычная концепция языка, которая позволяет отложенно выполнять блоки кода: например, чтобы закрывать файлы по завершении работы с ними.
Можно рассматривать defer как замену деструкторов/менеджеров контекста в других языках (try_with_resources из Java, with из Python).
func foo() {
    // паникуем
    panic("unexpected!")
}
//...
    // выполняется после срабатывания паники
    defer func() {
        if r := recover(); r != nil {
            // обработка паники, в переменной r будет лежать строка "unexpected"
        }
    }()
    // внутри foo срабатывает паника
    foo()


    Go даёт возможность перехватывать и обрабатывать панику. Для этого используется конструкция defer и встроенная функция recover.
    defer — это ещё одна необычная концепция языка, которая выполняет блоки кода при выходе из функции, например, чтобы закрывать файлы по завершении работы с ними.

    recover — функция, которая позволяет восстановить выполнение программы в случае паники.
*/
/*
swap

var a, b int
a, b = 5, 10 // a == 5, b == 10
a, b = b, a // swap: a == 10, b == 5

*/

/*
Строки в Go представляют собой массив из значений типа byte.
 По этой причине к элементам строки можно обращаться по индексу, а к самим строкам применима встроенная функция len, которая возвращает её длину в байтах:

 var a string
a = "abc"
println(len(a)) // 3
*/
