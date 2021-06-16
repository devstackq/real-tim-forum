import Parent from "./Parent.js";

export default class Profile extends Parent {
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
            let result = await response.json()
                console.log(result, 'res')
            } else {
                console.log('not uuid || incorrect')
                window.location.replace('/signin')
            }
    }

    async getHtml() {
        return super.showHeader('auth');
    }
}