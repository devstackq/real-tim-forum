import { redirect } from "../index.js";
// import { getUserId } from "./WebSocket.js";
let singletonInstance = null;

export default class Parent {
  constructor(text, type) {
    this.text = text;
    this.type = type;
    this.item = [];
    this.userId = 0;
    this.session = "";
    this.isAuth = "";
    this.auth = { state: false };
    // this.category =  document.cookie.split(";")[2].slice(11)
    this.vote = {
      id: 0,
      creatorid: 0,
      type: "",
      countlike: 0,
      countdislike: 0,
      group: "",
    };
    // if (!singletonInstance) {
    //   singletonInstance = this;
    // }
    // return singletonInstance;
  }

  setPostParams(group, id) {
    this.vote.group = group;
    this.vote.id = id;
  }

  getUserId() {
    if (document.cookie.split(";").length > 1) {
      return (this.userId = document.cookie.split(";")[1].slice(9).toString());
    }
  }

  getUserSession() {
    if (document.cookie.split(";").length > 1) {
      return (this.session = document.cookie.split(";")[0].slice(8).toString());
    }
  }

  setAuthState() {
    this.auth.state = true;
    this.isAuth = "true";
  }
  getAuthState() {
    if (document.cookie.split(";").length > 1) {
      return (this.isAuth = localStorage.getItem("isAuth"));
    }
  }

  getLocalStorageState(type) {
    return localStorage.getItem(type);
  }

  async fetch(endpoint, object) {
    let response = await fetch(`http://localhost:6969/api/${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify(object),
    });
    if (response.status == 200) {
      let result = await response.json();
      return result;
    } else {
      return null;
      // return response.statusText;
    }
  }

  renderSequence(object, ...type) {
    if (object != null) {
      if (type == "#postById") {
        this.render([object], `${type}`, "");
      }
      if (object.User != null) {
        this.render([object.User], ".bioUser", "User data ");
      }
      if (object.Posts != null) {
        this.render(object.Posts, ".postsUser", "Created posts");
      }
      if (object.VotedItems != null) {
        this.render(object.VotedItems, ".votedPost", "Voted posts");
      }
      if (object != null && type == "posts") {
        let category = "";
        if (this.isAuth == "true") {
          //use regex?
          category = document.cookie.split(";")[2].slice(11);
        } else if (this.isAuth == "false") {
          category = document.cookie.split(";")[0].slice(10);
        }
        this.render(object, ".postContainer", `${category} posts`);
      }
    }
  }

  render(seq, where, text) {
    let parent = document.querySelector(where);
    let title = document.createElement("p");
    parent.append(title);

    seq.forEach((item) => {
      title.textContent = text;
      let div = document.createElement("div");
      for (let [i, v] of Object.entries(item)) {
        let span = document.createElement("span");
        //case post by id
        i == "countlike" ? (span.id = "countlike") : "";
        i == "countdislike" ? (span.id = "countdislike") : "";
        span.textContent = ` ${i} : ${v} `;
        div.append(span);
      }
      //case User data & postById - not onlclick
      if (!item["email"] && where != "#postById") {
        div.value = item["id"];
        div.onclick = () => redirect(`postget?id=${item["id"]}`);
      }
      parent.append(div);
    });
  }

  createElement(...params) {
    // console.log(params[0]);
    let x = null;
    for (let [k, v] of Object.entries(params[0])) {
      if (v["type"] != undefined) {
        x = document.createElement(v["type"]);
      }
      if (v["id"] != undefined) {
        x.id = v["id"];
      }
      if (v["text"] != undefined) {
        x.textContent = v["text"];
      }
      if (v["value"] != undefined) {
        x.value = v["value"];
      }
      if (v["class"] != undefined) {
        x.className = v["class"];
      }
      if (v["attr"] != undefined) {
        x.setAttribute("id", v["id"]);
      }
      if (v["child"] != undefined) {
        x.appendChild(v.child);
      }
      if (v["parent"] != undefined) {
        v["parent"].append(x);
      }
    }
  }

  fillObject(obj) {
    for (let [k, v] of Object.entries(obj)) {
      if (document.getElementById(k) != null) {
        obj[k] = document.getElementById(k).value;
      }
    }
    if (obj["creatorid"] != null) {
      obj["creatorid"] = this.userId;
    }
    return obj;
  }

  async voteItem() {
    this.vote.creatorid = this.getUserId();
    this.vote.countdislike = document.querySelector("#countdislike").value;
    this.vote.countlike = document.querySelector("#countlike").value;

    let object = await this.fetch("vote", this.vote);

    if (object != null) {
      document.querySelector(
        "#countlike"
      ).textContent = ` countlike: ${object["countlike"]} `;
      document.querySelector(
        "#countdislike"
      ).textContent = `countdislike: ${object["countdislike"]} `;
    } else {
      console.log(object, "vote");
      // window.location.replace("/signin");
    }
  }

  voteDislike() {
    document.getElementById("btndislike").onclick = async () => {
      this.vote.type = "dislike";
      this.voteItem();
    };
  }
  voteLike() {
    document.getElementById("btnlike").onclick = async () => {
      this.vote.type = "like";
      this.voteItem();
    };
  }

  showHeader() {
    this.getUserId();
    this.isAuth = localStorage.getItem("isAuth");
    // this.getAuthState()
    let login = "";
    let register = "";
    let logout = "";
    let profile = "";
    let chat = "";

    if (this.isAuth == "false" || this.isAuth == null) {
      profile = "";
      logout = "";
      chat = "";
      register = `<a href="/signup"  class="nav__link signup" data-link>Signup</a>`;
      login = `<a href="/signin"  class="nav__link signin" data-link>Signin</a>`;
    } else if (this.isAuth == "true") {
      register = "";
      login = "";
      logout = `<a href="/logout"  id='logout' class="nav__link logout" data-link>Logout</a>`;
      profile = `<a href="/profile" class="nav__link" data-link>Profile</a>`;
      chat = `<a href="/chat" class="nav__link" data-link>Chat</a>`;
    }
    return `
        <nav class="nav">
        <a href="/all" class="nav__link" data-link>Main</a>
        ${login}
        ${register}
        <div class="dropdown">
          <button class="dropbtn">Categories</button>
          <div class="dropdown-content">
          <a href="/love" data-link>love</a>
          <a href="/science" data-link>science</a>
          <a href="/nature" data-link>nature</a>
        </div>
        </div>
        ${chat}
       ${profile}
        ${logout}
    </nav>
    <span id='notify' class='notify' > </span> 
`;
  }

  showNotify(text, type) {
    let notify = document.getElementsByClassName("notify")[0];
    if (type == "error") {
      notify.style.display = "block";
      notify.textContent = text;
    } else if (type == "hide") {
      notify.style.display = "none";
    }
  }
}
