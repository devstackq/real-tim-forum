import Parent from "./Parent.js";

export default class ViewPost extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }
  setTitle(title) {
    document.title = title;
  }

  async postById(id) {
    let postId = { id: 0 };
    postId.id = id;

    let object = await super.fetch("post/id", postId);

    if (object != null) {
      // let status, res  = this.Fetch('post/id', postId)
      let parent = document.querySelector(".postParent");
      parent.innerHTML = "";

      let btnDislike = document.createElement("button");
      let btnTextarea = document.createElement("button");

      //   let btnLike = document.createElement("button");
      //   btnLike.id = "btnLike";
      //   btnLike.textContent = "like";
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
      //   parent.append(btnLike);
      parent.append(btnDislike);
      let auth = localStorage.getItem("isAuth");
      // let auth = this.getLocalStorageState('isAuth')
      if (auth == "true") {
        parent.append(textarea);
        parent.append(btnTextarea);
        super.showHeader("auth");
      }

      super.createElement([
        { type: "button" },
        { id: "btnLike" },
        { text: "like" },
        { parent: parent },
        { func: super.votePost },
      ]);

      //lost comment here ?
      btnTextarea.onclick = async () => {
        let text = document.getElementById("commentField").value;
        let result = await this.fetch("comment", text);
        console.log(
          result,
          "get data from text area, then create comment by post id, then -> show under current post new comment"
        );
      };
      //   btnLike.onclick = () => {
      //     this.votePost("like");
      //   };
      btnDislike.onclick = () => {
        super.votePost("dislike");
      };
    }
  }


  async init() {
  
    let id = window.location.href.split("=")[1]
    // let id  = url.searchParams.get("id")
    console.log(id)
    let result = this.postById(id);
    if (result != null) {
      console.log(result);
    } else {
      console.log("error logout");
      super.showNotify("not found post", "error");
    }
  }

  async getHtml() {
    createLikeButton etc
    let body = `
    
        <div id='postParent'> </div>
    `;
    return super.showHeader("auth") + body;
  }
}
