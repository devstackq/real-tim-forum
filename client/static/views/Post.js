import Parent from "./Parent.js";

export default class ViewPost extends Parent {
    constructor(params) {
        super();
        this.params = params;
        this.postId = window.location.href.split("=")[1];
        this.vote = {
            id: window.location.href.split("=")[1],
            creatorid: super.getUserId(),
            type: "",
            group: "post"
        }
    }
    setTitle(title) {
        document.title = title;
    }

    async postById(id) {

        let object = await super.fetch("post/id", this.vote);
        let parent = document.querySelector("#postParent");
        if (object != null) {

            let btnTextarea = document.createElement("button");
            btnTextarea.textContent = "lost comment";

            for (let [k, v] of Object.entries(object)) {
                let span = document.createElement("span");
                if (v != null && v != "") {
                    span.textContent = `${k} : ${v}`;
                }
                parent.append(span);
            }
        } else {
            super.showNotify('bad request', 'error')
            parent.innerHTML = ""
        }
    }

    async postVote() {
            console.log(this.vote, "param vote")
            let object = await super.fetch("post/vote", this.vote);
            if (object != null) {
                //like postById, update page ?
                window.location.reload
            } else {
                window.location.replace('/signin')
            }
        }
        //preload, then wait event
    postDislike() {
        document.getElementById("btndislike").onclick = async() => {
            this.vote.type = "dislike";
            this.postVote()
        };
    }
    postLike() {
        document.getElementById("btnlike").onclick = async() => {
            this.vote.type = "like";
            this.postVote()
        }
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
        super.createElement([
            { type: "button" },
            { id: "btnlike" },
            { text: "like" },
            { parent: parent },
            { value: 'like' },
            { func: this.postLike },
        ]);

        super.createElement([
            { type: "button" },
            { id: "btndislike" },
            { text: "dislike" },
            { parent: parent },
            { func: this.postDislike },
        ]);
        //init event onclcik
        this.postLike()
        this.postDislike()
    }

    async getHtml() {

        let body = `
        <div id='postParent'>  </div>
    `;
        let comment = `<textarea id="commentField"> </textarea>  <button id="btncomment">lost comment </button>`

        let authState = localStorage.getItem("isAuth");

        if (authState == "true") {
            return super.showHeader("auth") + body + comment
        } else {
            return super.showHeader("free") + body
        }
    }
}