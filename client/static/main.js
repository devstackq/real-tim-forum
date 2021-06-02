import Main from "./views/MainPage.js";
import Profile from "./views/Profile.js";
import Signup from './views/Signup.js';

const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const getParams = match => {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);
    console.log("get param", keys, values)

    return Object.fromEntries(keys.map((key, i) => {
        return [key, values[i]];
    }));
};

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
};

const signup = ()=> {
        
         document.getElementById('signup').onclick = async function() {
        
            let e = document.getElementById("email").value;
            let p = document.getElementById("password").value;
        
            let user = {
                email: e,
                password: p
            };
            console.log(user, "user")
            let response = await fetch('http://localhost:6969/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            });
            let result = await response.json();    
console.log(result, 'result')        
        }
}

const router = async() => {
    const routes = [
        { path: "/", view: Main },
        { path: "/profile", view: Profile },
        { path: "/signup", view : Signup},
        // { path: "/post/:id", view: PostView },
    ];

    // Test each route for potential match
    const potentialMatches = routes.map(route => {
        return {
            route: route,
            result: location.pathname.match(pathToRegex(route.path))
        };
    });

    let match = potentialMatches.find(potentialMatch => potentialMatch.result !== null);

    if (!match) {
        match = {
            route: routes[0],
            result: [location.pathname]
        };
    }

    const view = new match.route.view(getParams(match));
    console.log(view, "view")

    document.querySelector("#app").innerHTML = await view.getHtml();

    if(document.getElementById('signup') !== null) {
        signup()
    }
};


window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });
    router();
});