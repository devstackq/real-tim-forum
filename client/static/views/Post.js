import Parent from "./Parent.js";

export default class ViewPost extends Parent {
  constructor(params) {
    super();
    this.params = params;
    this.postId = window.location.href.split("=")[1];
  }
  setTitle(title) {
    document.title = title;
  }

  async postById(id) {
    let postId = { id: "" };
    postId.id = id;

    let object = await super.fetch("post/id", postId);

    if (object != null) {
      let parent = document.querySelector("#postParent");

      let btnTextarea = document.createElement("button");
      btnTextarea.textContent = "lost comment";

      for (let [k, v] of Object.entries(object)) {
        let span = document.createElement("span");
        if (v != null && v != "") {
          span.textContent = `${k} : ${v}`;
        }
        parent.append(span);
      }
    }
  }

//preload, then wait event
  votePost() {
    
    let vote = { type: "", postid: 0, userid: 0 };

    document.getElementById("btndislike").onclick = async () => {

      vote.type = "dislike";
      vote.postid = this.postId;
      vote.userid = super.getUserId();

      //
      let object = await super.fetch("vote", vote);
      if (object != null) {
        //like postById
        console.log("vote state", object);
      }
    };
    
    document.getElementById("btnlike").onclick = async () => {

        vote.type = "like";
        vote.postid = this.postId;
        vote.userid = super.getUserId();
  
        let object = await super.fetch("vote", vote);
        if (object != null) {
          //like postById
          console.log("vote state like", object);
        }else {
            super.showNotify('vote can work now', 'error')
        }
      };
  }

  async lostComment() {
    let content = document.getElementById("commentField").value;

    let comment = { content: "", userid: 0, postid: 0 };
    comment.content = content;
    // comment.postId =
    // comment.uid =

    let object = await super.fetch("comment", comment);
    if (object != null) {
      //like postById
      console.log("lost comment in post, who ?", object);
    }
  }

  async init() {
    let parent = document.querySelector("#postParent");
    // let id  = url.searchParams.get("id")

    let result = this.postById(this.postId);

    if (result != null) {
      console.log(result);
    } else {
      super.showNotify("not found post", "error");
    }
    // {parent : document.getElementById('commentField')},
    super.createElement([
      { type: "textarea" },
      { id: "commentField" },
      { parent: parent },
    ]);
    super.createElement([
      { type: "button" },
      { id: "btncomment" },
      { text: "lost comment" },
      { parent: parent },
      { func: this.lostComment },
    ]);

    super.createElement([
      { type: "button" },
      { id: "btnlike" },
      { text: "like" },
      { parent: parent },
      {value : 'like'},
      { func: this.votePost },
    ]);

    super.createElement([
      { type: "button" },
      { id: "btndislike" },
      { text: "dislike" },
      { parent: parent },
      { func: this.votePost },
    ]);

    this.votePost();
  }

  async getHtml() {
    // createLikeButton etc
    let body = `
        <div id='postParent'> </div>
    `;
    return super.showHeader("auth") + body;
  }
}
