# Altair Technical Test

Technical test in which a list is separated into two lists, ReadTeam and BlueTeam, depending on the teams to which the people from the initial list belongs.

For the red team, given an id, it will be possible to know if the user's word corresponding to that id is a palindrome or not.

For the blue team, given an id, it will be possible to know if the user's word corresponding to that id contains five vowels.

## Dependencies

As the main language of the application you must have Go installed on the device where you are going to run it.

## Example

In order to run the application you should go into <strong>cmd</strong> directory and execute <strong>go run main.go</strong>.

Once this is done there will be an API listening on port 8080.


### Endpoints

There are two differents endpoints

<ol>
  <li>/teams/vowel-filter/{id}
      <ol><br> 
      <li>Returns <strong>200 - true</strong> in case the user corresponding to the given id is in the BlueTeam and its associated word has 5 vowels.</li><br> 
      <li>Returns <strong>200 - false</strong> in case the user corresponding to the given id is in the BlueTeam and its associated word does not have 5 vowels.</li><br> 
       <li> Returns <strong>404 - NotFound</strong> in case the user corresponding to the given id does not exist in the BlueTeam..</li><br> 
    </ol><br> 
  <li>/teams/palindrome-filter/{id}
     <ol><br> 
      <li>Returns <strong>200 - true</strong> in case the user corresponding to the given id is in the RedTeam and its associated word is palindrome.</li><br> 
      <li>Returns <strong>200 - false</strong> in case the user corresponding to the given id is in the RedTeam and its associated word is not palindrome.</li><br> 
     <li> Returns <strong>404 - NotFound</strong> in case the user corresponding to the given id does not exist in the RedTeam..</li><br> 
    </ol>
</ol>

### Data

Data used in this technical test

       personas = [{  
            "id":"AAA",

            "nombre":"Javier",

            "apellido":"Miguez",

            "equipo":"rojo",

            "edad":44,

            "palabra":"riachuelo"

            },

            {

            "id":"AAB",

            "nombre":"Luis",

            "apellido":"Alonso",

            "equipo":"azul",

            "edad":32,

            "palabra":"murcielago"

            },

            {

            "id":"ABA",

            "nombre":"Cristina",

            "apellido":"Perez",

            "equipo":"azul",

            "edad":21,

            "palabra":"solos"

            },

            {

            "id":"ABB",

            "nombre":"Mariano",

            "apellido":"Gomez",

            "equipo":"rojo",

            "edad":18,

            "palabra":"reconocer"

            },

            {

            "id":"BAA",

            "nombre":"Alicia",

            "apellido":"Domenech",

            "equipo":"rojo",

            "edad":32,

            "palabra":"oso"

            },

            {

            "id":"BBA",

            "nombre":"Susana",

            "apellido":"Gomez",

            "equipo":"azul",

            "edad":19,

            "palabra":"elefante"

            }

    ]




