export default class {
    constructor(params) {
        this.params = params;
    }
    setTitle(title) {
        document.title = title;
    }
    async init() {
        //get by user_id -> data(posts/comments, votes) from backend
        let response = await fetch('http://localhost:6969/api/profile')
        let result = await response.json();
        console.log(result, ', profile page func js')
    }
    async getHtml() {
        //get -> set profile data -> by id
        //fetch - post, edit -> name etc
        return "profile page";
    }
}