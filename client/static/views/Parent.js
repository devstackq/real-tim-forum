export default class Parent {
  constructor(text, type) {
    this.text = text;
    this.type = type;
    this.item = [];
    this.userId = 0;
    this.vote = {
      id: 0,
      creatorid: 0,
      type: "",
      group: "",
    };
  }
  setPostParams(group, id) {
    this.vote.group = group;
    this.vote.id = id;
  }

  getUserId() {
    return this.userId.toString();
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

  createElement(...params) {
    console.log(params[0]);
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
      //check funcType == 'vote', comment
      if (v["func"] != undefined) {
        console.log(v["text"]);
        x.onclick = () => {
          v.func("like");
        };
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
  // create render uniq func DRY

  //   async postById(post) {

  //     let object = await fetch("post/id", post);

  //     let parent = document.querySelector("#postParent");
  //     if (object != null) {
  //       let btnTextarea = document.createElement("button");
  //       btnTextarea.textContent = "lost comment";
  //       console.log(post, 1)

  //       for (let [k, v] of Object.entries(object)) {
  //         let span = document.createElement("span");
  //     if(k =='countlike') {
  //         span.id = 'countlike'
  //     }
  //     if(k =='countdislike') {
  //         span.id = 'countdislike'
  //     }
  //         if (v != null) {
  //             span.textContent = ` ${k} : ${v} \n`;
  //         }
  //             parent.append(span);
  //       }
  //     } else {
  //       showNotify("bad request", "error");
  //       // parent.innerHTML = ""
  //     }
  //   }

  async voteItem() {
    this.vote.creatorid = this.getUserId();

    let object = await this.fetch("vote", this.vote);
    if (object != null) {
      console.log(object, "vote data", object['id']);
      //DRY update value by key or new show postById
      document.querySelector("#countlike").textContent = ` countlike: ${object["countlike"]} `
      document.querySelector("#countdislike").textContent = `countdislike: ${object["countdislike"]} `
    } else {
      window.location.replace("/signin");
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

  showHeader(type) {
    if (document.cookie.split(";").length > 1) {
      this.userId = document.cookie.split(";")[1].slice(9);
    }

    let login = "";
    let register = "";
    let logout = "";
    let profile = "";

    if (type == "free") {
      profile = "";
      logout = "";
      register = `<a href="/signup"  class="nav__link signup" data-link>Signup</a>`;
      login = `<a href="/signin"  class="nav__link signin" data-link>Signin</a>`;
    } else if (type == "auth") {
      register = "";
      login = "";
      logout = `<a href="/logout"  id='logout' class="nav__link logout" data-link>Logout</a>`;
      profile = `<a href="/profile" class="nav__link" data-link>Profile</a>`;
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
       ${profile}
        ${logout}
    </nav>
    <span class='notify' > </span> 
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
