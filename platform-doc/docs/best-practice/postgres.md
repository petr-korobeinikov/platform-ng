# Postgres

## Ведите разработку приложения в выделенной схеме

Создайте отдельную схему для приложения:

```postgresql
create schema app;
```

Это позволит избежать конфликтов с функциями и типами данных из устанавливаемых
расширений.

## Устанавливайте расширения в отдельные схемы

```postgresql
create extension "uuid-ossp" with schema uuid;
create extension intarray with schema intarray;
```

## Заведите отдельную схему `maintenance` для служебных данных

```postgresql
create schema maintenance;
```

В ней удобно хранить

* данные о применённых миграциях;
* таблицы, необходимые для временного хранения данных при их переносе;

## Используйте единственное число

<!-- @formatter:off -->
```postgresql
create table person ();
create table item ();
create table request ();
```
<!-- @formatter:on -->

Заводя таблицы в `postgres`, Вы, фактически, создаёете тип данных.

```postgresql
create table person
(
    name    text,
    surname text
);

select ($$(John,Doe)$$::person).name;
```

Вы же не именуете типы данных во множественном числе?

```go
type Requests struct { ... }

type People struct {
Name, Surname string
}
```

## Не используйте `if not exists`

Чаще всего конструкцию `if not exists` используют в миграциях для обхода
конфликтных ситуаций.

Наличие конструкции `if not exists` сигнализирует о проблемах в пайплайне
применения миграций:

* отсутствует должный контроль над применением миграций;
* отсутствует трекинг применённых изменений.

Предаставим ситуацию, в которой `if not exists` создаёт большую проблему, чем
попытка обеспечить идемпотентность изменений.

В базе данных создана таблица вида:

```postgresql
create table person
(
    id      bigserial primary key,
    name    text,
    surname text
);
```

Миграция содержит следующие изменения:

```postgresql
create type person_initial as
(
    name    text,
    surname text
);

create table if not exists person
(
    id         uuid primary key,
    created_at timestamp      not null,
    initial    person_initial not null
);
```

В итоге мы имеем две совершенно разные таблицы, к определению которых код
приложения не готов.

## Используйте `text` вместо `varchar`

1. `text` короче, чем `varchar`.
2. `text` принято использовать в `postgresql`-сообществе.
3. `text` и `varchar` используют один и тот же механизм хранения и тип
   данный `varlena`.
4. Нет особого смысла навешивать ограничение длины строки на базе в
   виде `varchar(16)`.

    * Во-первых, проверка на стороне приложения проще и быстрее.
    * Во-вторых, обновление типа для больших таблиц при расширении диапазона
      потребует эксклюзивной блокировки, что может вызвать длительный простой
      приложения.

## Пишите sql в нижнем регистре

А стандарт же? Как же стандарт?

Выразительные `sql`-запросы отлично пишутся и отлично читаются даже при
форматировании в нижнем регистре.

В `postgresql`-сообществе "капсить" считается моветоном.

## Не используйте `*`

* Всегда явно перечисляйте имена колонок:
    * в выражениях `select ...`,
    * в выражениях `returning ...`,
    * при вставке в выражениях `insert into t (...)`.

## Полезные ссылки

* Отличный материал, полностью совпадающий с мнением
  автора: [https://levelup.gitconnected.com/how-to-design-a-clean-database-2c7158114e2f](https://levelup.gitconnected.com/how-to-design-a-clean-database-2c7158114e2f)
