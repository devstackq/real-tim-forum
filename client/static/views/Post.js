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
    super.voteLike("#countlike", "btnlike");
    super.voteDislike("#countdislike", "btndislike");

    //out -> Parent -> use Comment & Post component
    lostCommentId.onclick = async () => {
      let content = document.getElementById("commentField").value;
      let comment = { content: "", creatorid: 0, postid: 0 };
      comment.content = content;
      comment.postid = parseInt(this.post.id);
      comment.creatorid = super.getCookie("user_id");
      if (comment.creatorid != undefined) {
        if (
          comment.content != 0 &&
          comment.content != "" &&
          comment.content.length > 0
        ) {
          let object = await super.fetch("comment", comment);
          if (object != null) {
            //show under comment
            let view = document.getElementById("comment_container");
            let div = document.createElement("div");
            div.textContent = `content : ${object.content} creatorid :  ${object.creatorid} createdtime : ${object.createdtime} author: ${object.author}`;
            // super.commentField(object.id);   //set last comment voteFunc()
            view.append(div);
            //append last comment in post
            document.getElementById("commentField").value = "";
          }
        } else {
          alert("empty value");
        }
      } else {
        window.location.replace("/signin");
      }
    };
  }

  async getHtml() {
    let body = `<div id='postById'> <div id=comment_container><textarea id="commentField"> </textarea>  <button  id="lostCommentId">lost comment</button> </div> </div>`;
    let comment = ``;
    return super.showHeader() + body + comment;
  }
}
