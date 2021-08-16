import Parent from "./Parent.js";

export default class ViewPost extends Parent {
  constructor(params) {
    super();
    this.params = params;
    this.post = {
      id: window.location.href.split("=")[1],
    };
    this.isAuth = localStorage.getItem("isAuth");
    this.comment = { id: 0 };
  }

  setTitle(title) {
    document.title = title;
  }

  async postById() {
    let object = await super.fetch("post/id", this.post);
    if (object != null) {
      super.renderSequence(object, "#postById");
    } else {
      super.showNotify("bad request", "error");
      // parent.innerHTML = ""
    }
  }
  async getCommentsByPostId(pid) {
    let obj = { postid: 0 };
    obj.postid = parseInt(pid);

    let response = await super.fetch("comment/id", obj);
    if (response != null) {
      //show under comment
      response.forEach((item) => {
        super.renderSequence(item, "#comment_container");
      });
    }
  }

  async init() {
    let parent = document.querySelector("#postById");
    // let id  = url.searchParams.get("id")
    this.postById(this.post.id);
    super.createElement([
      { type: "button" },
      { id: "btnlike" },
      { text: "like" },
      { parent: parent },
      // { value: "like" },
    ]);

    super.createElement([
      { type: "button" },
      { id: "btndislike" },
      { text: "dislike" },
      { parent: parent },
      // { func: this.postDislike },
    ]);
    //init event onclick
    super.setPostParams("post", this.post.id);
    super.voteLike("btnlike");
    super.voteDislike("btndislike");

    this.getCommentsByPostId(this.post.id); // get all comment by post

    //out -> Parent -> use Comment & Post component
    lostCommentId.onclick = async () => {
      let content = document.getElementById("commentField").value;
      let comment = { content: "", creatorid: 0, postid: 0 };
      comment.content = content;
      comment.postid = parseInt(this.post.id);
      comment.creatorid = super.getCookie("user_id");
      console.log("send quey vote comment");

      let object = await super.fetch("comment", comment);
      if (object != null) {
        //show under comment
        let view = document.getElementById("comment_container");
        let div = document.createElement("div");
        div.textContent = `${object.id}, ${object.content}, ${object.countlike}, ${object.countdislike}`;
        super.commentField(object.id);
        view.append(div);
        // this.comment = object.id;
        console.log("lost comment in post, who ?", object);
        //append last comment in post
      }
    };
  }

  async getHtml() {
    //?DRY
    let body = `<div id='postById'> <div id=comment_container> </div> </div>`;
    let comment = `<textarea id="commentField"> </textarea>  <button  id="lostCommentId">lost comment</button>`;
    return super.showHeader() + body + comment;
  }
}
