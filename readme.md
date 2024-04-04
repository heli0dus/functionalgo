<img src="./fungologo.png" alt="drawing" width="800"/>

# FunGo

FunGo - библиотека для реализации возможностей функционального программирования в языке go. Состоит из следующих пакетов

- **collections** - предоставляет некотоыре коллекции с функциональными инетрфейсами
    - collections/arrayList 
    - collections/hashSet
    - collections/maybe
    - collections/linkedList
    - collections/forwardList
- **functools** - предоставляет функции высшего порядка для композиции и продвинутой работы с базовыми функциями
- **stream** - представляет безовые функции по аналогии с Java Stream API, а именно
    - `Fmap(f)` - применяет переданную функцию ко всем элементам стрима
    - `FlatMap(f)` - применяет переданную функцию возвращающую срез и объединяет полученные срезы в один общий стрим
    - `Take(x)` - берёт первые `x` элементов стрима
    - `TakeWhile(p)` - сохраняет в стриме первые элементы пока они соответствуют предикату `p`
    - `Drop(x)` - отбрасывает первые `x` элементов стрима
    - `DropWhile(p)` - отбрасывает элементы стрима пока они соответствуют предикату `p`
    - `Filter(p)` - сохраняет в стриме только элементы, при которых выполянется предикат `p`
    - `Reduce(f, init)` - левая свёртка стрима
    - `All(p)` - `True`, если для всех элементов стрима выполняется предикат `p`
    - `Any(p)` - `True`, если хотя бы для одного элемента стрима выполняется предикат `p`

# Технические детали

Для коллекций функции реализованы при помощи generic функций. Стримы же используют типовую рефлексию

# Пример

В файле `main.go` приведен пример использования библиотеки стримов. Команда для сборки и запуска примера:

```bash
go run src/main.go
```