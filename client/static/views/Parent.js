export default class Parent {
  constructor(text, type) {
    this.text = text;
    this.type = type;
    this.item = [];
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
    }
  }

  async votePost(type) {
    console.log(type, 2);
    let vote = { type: "" };
    vote.type = type;

    let object = await this.fetch("vote", vote);
    if (object != null) {
      console.log("vote state", object);
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
      //check funcType == 'vote', comment
      if (v["func"] != undefined) {
        console.log(v["text"]);
        x.onclick = () => {
          v.func("like");
        };
      }
    }
  }

  // send func inside -> getPost component ?
  
  //isAuth - middleware -> showHeader - change
  // creatEleme func todo
  // another url path ->  server  error fix
  // fix - > each post -> own url/id
// create func - render - for uniq type, where etc
  
fillObject(obj) {
    for (let [k, v] of Object.entries(obj)) {
      if (document.getElementById(k) != null) {
        obj[k] = document.getElementById(k).value;
      }
    }
    return obj;
  }

  //render -> send component inside  Posts
  
  render(item, idx, where, type) {
    let wrapper = document.querySelector(where);
    let btn = document.createElement("button");
    let div = document.createElement("div");
    // let className=''
    // let f = null
    //     if(type =='post') {
    //         className = 'postWrapper'
    //         f = this.postById
    //     }else if (type=='profile'){
    //         className = 'profileWrapper'
    //         f = this.editProfile
    div.className = "postWrapper";

    for (let [k, v] of Object.entries(item)) {
      if (v != "" && v != null) {
        let span = document.createElement("span");
        span.id = k;
        span.textContent = ` ${k} : ${v}`;
        btn.value = idx;
        btn.textContent = `see post`;

        btn.onclick = () => {
        window.location.replace(`/postget?id=${item["id"]}`)
          // this.postById(item["id"]);
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
