# rucksack

A rucksack is distributed key value database.
The project is essentially a result of an attempt to get deeper into dystributed systems and database internals.
It's inspired by [tinyKV course](https://github.com/talent-plan/tinykv) and [MIT 6.5840](https://pdos.csail.mit.edu/6.824/index.html).
Overall architechture, some design desicions and parts of code will be taken from [tinyKV course](https://github.com/talent-plan/tinykv), with copyrights of course.

__Why not to re-use tinyKV repo?__
TinyKV contains lots of things that confuse me. Since the main goal is to build a DB, but also
undertand approaches and tradeoffs, I think it's better to start from scratch.