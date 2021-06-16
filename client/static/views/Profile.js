import Parent from "./Parent.js";

export default class Profile extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }
    setTitle(title) {
        document.title = title;
    }
    //auto create or create then -> fill data ?
    showBio(data) {
        let div = document.createElement('div')
         for( let [k,v] of Object.entries( data)) {
            console.log(k,v)
            let span = document.createElement('span')
            span.textContent = `${k} : ${v}`
            div.append(span)    
        }
        document.body.append(div)
    }

    async init() {
        //get by user_id -> data(posts/comments, votes) from backend
        let response = await fetch('http://localhost:6969/api/profile')
        // console.log(response.status, response)
        if (response.status === 200) {
            let result = await response.json()
        this.showBio(result)
            // console.log(result, 'res')
            //show data
            } else {
                console.log('not uuid || incorrect')
                window.location.replace('/signin')
        }
    }

    async getHtml() {

        return super.showHeader('auth');
    }
}