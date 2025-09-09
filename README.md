# real-time-forum

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-07405E?style=flat-square&logo=sqlite&logoColor=white)](https://www.sqlite.org/)
[![WebSocket](https://img.shields.io/badge/WebSocket-000000?style=flat-square&logo=websocket&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/API/WebSockets_API)
[![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=flat-square&logo=html5&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/HTML)
[![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=flat-square&logo=css3&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/CSS)
[![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=flat-square&logo=javascript&logoColor=black)](https://developer.mozilla.org/fr/docs/Web/JavaScript)
[![Font Awesome](https://img.shields.io/badge/Font_Awesome-528DD7?style=flat-square&logo=font-awesome&logoColor=white)](https://fontawesome.com/)

Un forum en temps réel permettant aux utilisateurs de communiquer et de voir en direct quand d'autres utilisateurs sont en train d'écrire.

## 🛠 Technologies Utilisées

### Backend
- **Langage** : Go 1.18
- **Base de données** : SQLite3
- **WebSockets** : gorilla/websocket
- **Gestion des sessions** : Sessions personnalisées avec UUID
- **Sécurité** : bcrypt pour le hachage des mots de passe
- **Gestion des dépendances** : Go Modules

### Frontend
- **HTML5** : Structure des pages
- **CSS3** : Styles et mises en page
- **JavaScript Vanilla** : Interactions client-side
- **Font Awesome** : Icônes
- **WebSockets** : Communication en temps réel avec le serveur

### Outils de développement
- **Système de contrôle de version** : Git
- **Gestion de base de données** : Fichiers SQL pour la structure

## Fonctionnalités

- Inscription et authentification des utilisateurs
- Chat en temps réel avec indication de frappe (typing indicator)
- Gestion des sessions utilisateur
- Interface utilisateur réactive et moderne

---

A typing in progress engine is a way that people can see that a user is typing in real time. Allowing you to see the other user is replying or sending a message.

The typing in progress engine must work in real time! This meaning that if you start typing to a certain user this user will be able to see that you are typing.

This engine must have/display:

    A websocket to establish the connection with both users
    An animation so that the user can see that you are typing, this animation should be smooth (no interruptions/janks) and just enough to draw attention for the user to see (user friendly)
    The name of the user that is typing
    Whenever the user stops typing or finishes the conversation, it should not display the animation

To help with the display of the typing in progress you can take a look on the js event list, mainly the Keyboard events and the Focus events

## How to use the Application

1. Clone this repository on your computer using the following command:
```bash
  git clone https://github.com/badStephane/real-time-forum.git
```

2. Go to the project directory:
```bash
  cd real-time-forum/
```

3. Run the application using Node.js with the following command:
```bash
  go run .
```


---

Connect, Share, Engage - Building Community Together!✨
