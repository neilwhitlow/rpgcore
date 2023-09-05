# RPG Core

A package to provide a basic, game-system-agnostic, library for creating
pen and paper RPG utilities for GMs. 

>*Note: As of fall 2023, this should be considered experimental. I have mostly coded 
in C# and Java over the decades, and this is my first playground repo attempting to
write a standalone Go package for consumptions, and to experiment with various methods
of unit testing ranging from various approaches with the standard library, to eventually 
trying third-party testing packages such as testify.*

## rpgcore/dice 
Routines for generating randomizations for RPGs, most often expressed 
in terms and types of dice to be used. RNG uses the 64bit Mersenne Twister pseudo 
random number generator from the package [mt19937](https://github.com/seehuhn/mt19937).

## rpgcore/units 
Currently just a few units of measure and weight conversions. 

>*Again, unit testing is a free-for-all at the moment, with a ridiculous amount of repetitive
tests intended for changing and learning*
