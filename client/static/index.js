import Posts from "./views/MainPage.js";
import Profile from "./views/Profile.js";
import Signup from "./views/Signup.js";
import Signin from "./views/Signin.js";
import Logout from "./views/Logout.js";
import CreatePost from "./views/CreatePost.js";
import ViewPost from "./views/Post.js";
import Chat from "./views/Chat.js";

export const redirect = (endpoint) => {
  console.log(endpoint, "redirect");
  history.pushState(null, "", `http://localhost:6969/${endpoint}`);
  window.addEventListener("popstate", router());
};

const pathToRegex = (path) =>
  new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const getParams = (match) => {
  const values = match.result.slice(1);
  const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(
    (result) => result[1]
  );
  return Object.fromEntries(
    keys.map((key, i) => {
      return [key, values[i]];
    })
  );
};

const navigateTo = (url) => {
  history.pushState(null, null, url);
  console.log(url, "url");
  //if url == '/chat

  router();
};

const router = async () => {
  const routes = [
    { path: "/all", view: Posts },
    { path: "/love", view: Posts },
    { path: "/science", view: Posts },
    { path: "/nature", view: Posts },
    { path: "/profile", view: Profile },
    { path: "/signup", view: Signup },
    { path: "/signin", view: Signin },
    { path: "/logout", view: Logout },
    { path: "/postcreate", view: CreatePost },
    { path: "/postget", view: ViewPost },
    { path: "/chat", view: Chat },
  ];

  // Test each route for potential match
  const potentialMatches = routes.map((route) => {
    return {
      route: route,
      result: location.pathname.match(pathToRegex(route.path)),
    };
  });

  let match = potentialMatches.find(
    (potentialMatch) => potentialMatch.result !== null
  );

  if (!match) {
    match = {
      route: routes[0],
      result: [location.pathname],
    };
  }
  const view = new match.route.view(getParams(match));
  view.setTitle(match.result[0]);
  document.querySelector("#app").innerHTML = await view.getHtml();
  // let date = new Date(Date.now() + 86400);expires=${date.toUTCString()
  document.cookie = `category=${match.result[0]}; path=/; sameSite: "Lax";`;
  view.init();
  // if (window.location.pathname == "/chat") {
  //   window.location.reload();
  // }
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("click", (e) => {
    if (e.target.matches("[data-link]")) {
      e.preventDefault();
      navigateTo(e.target.href);
    }
  });
  router();
});
