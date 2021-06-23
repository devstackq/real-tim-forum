import Parent from "./Parent.js";

export default class ViewPost extends Parent {
    constructor(params) {
        super();
        this.params = params;
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

    async init() {

    }
    async getHtml() {
        //isAuth
        return super.showHeader("free")
    }
}