import Parent from "./Parent.js";

export default class extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }
    setTitle(title) {
        document.title = title;
    }

    async init() {
        //get by user_id -> data(posts/comments, votes) from backend
        let response = await fetch('http://localhost:6969/api/profile')

        if (response.status === 200) {
            // this.toggle()
        } else {
            console.log('not uuid || incorrect')
            window.location.replace('/signin')
        }
    }
    async getHtml() {
        //get -> set profile data -> by id
        //fetch - post, edit -> name etc
        let html = super.showHeader('auth');
        console.log(html, 'hed')
            // return h + wrapper
        return html
    }
}