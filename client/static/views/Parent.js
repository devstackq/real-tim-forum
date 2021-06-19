export default class Parent {
  constructor(text, type) {
    this.text = text;
    this.type = type;
    this.value = "";
    this.item = [];
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
      if (v["class"] != undefined) {
        x.className = v["class"];
      }
      if (v["child"] != undefined) {
        x.appendChild(v.child);
      }
      if (v["parent"] != undefined) {
        v["parent"].append(x);
      }
    }
  }

  async votePost(type) {
    let vote = { type: "" };
    vote.type = type;

    let object = await this.fetch("vote", vote);

    if (object != null) {
      console.log("vote state", object);
    }
  }

  async postById(id) {
    let postId = { id: 0 };
    postId.id = id;

    let object = await this.fetch("post/id", postId);

    if (object != null) {
      // let status, res  = this.Fetch('post/id', postId)
      let parent = document.querySelector(".postContainer");

      parent.innerHTML = "";

      let btnLike = document.createElement("button");
      let btnDislike = document.createElement("button");
      let btnTextarea = document.createElement("button");
      btnLike.id = "btnLike";

      btnLike.textContent = "like";
      btnDislike.textContent = "dislike";
      btnTextarea.textContent = "lost comment";

      let textarea = document.createElement("textarea");
      textarea.id = "commentField";

      for (let [k, v] of Object.entries(object)) {
        let span = document.createElement("span");
        if (v != null && v != "") {
          span.textContent = `${k} : ${v}`;
        }
        parent.append(span);
      }
      parent.append(btnLike);
      parent.append(btnDislike);
      parent.append(textarea);
      parent.append(btnTextarea);

    //   this.createElement([
    //     { type: "input" },
    //     { id: "maestro" },
    //     { text: "skr" },
    //     { child: btnLike },
    //     { parent: parent },
    //   ]);

      btnTextarea.onclick = () => {
        // this.votePost("like");
        console.log(
          "get data from text area, then create comment by post id, then -> show under current post new comment"
        );
      };

      btnLike.onclick = () => {
        this.votePost("like");
      };
      btnDislike.onclick = () => {
        this.votePost("dislike");
      };
    }
  }

  //render posts & getPostById
  render(item, idx, where) {
    let wrapper = document.querySelector(where);
    let btn = document.createElement("button");
    let div = document.createElement("div");

    for (let [k, v] of Object.entries(item)) {
      if (v != "" && v != null) {
        let span = document.createElement("span");
        div.className = "postWrapper";
        span.id = k;
        span.textContent = ` ${k} : ${v}`;

        btn.value = idx;
        btn.textContent = `see post`;

        btn.onclick = () => {
          this.postById(item["id"]);
        };
        div.appendChild(span);
        div.appendChild(btn);
        wrapper.appendChild(div);
      }
    }
  }

  showHeader(type) {
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
      logout = `<a href="/logout" id='logout' class="nav__link logout" data-link>Logout</a>`;
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
