import Parent from "./Parent.js";

export default class Posts extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }
  //async - await -> then hell
  async init() {
    let response = await fetch(`http://localhost:6969/api/post/`);
    if (response.status == 200) {
      let result = await response.json();
      //show all post
      result.forEach((element, idx) => {
        super.render(element, idx, ".postContainer");
      });
    }
  }

  async getHtml() {
    // let uuid = document.cookie.split(";")[1].slice(9, )
    let authState = localStorage.getItem("isAuth");
    let body = `<div class="postContainer"> </div>`;

    let create = `<a href="/postcreate" class="nav__link" data-link>create post</a>`
    console.log(authState)
    if (authState == "true") {
      return super.showHeader("auth")  + create + body 
    } else {
      return super.showHeader("free") + body + '';
    }
  }
}
