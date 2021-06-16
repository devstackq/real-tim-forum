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

        document.querySelector('#logout').onclick = async function() {

                let response = await fetch('http://localhost:6969/logout')
                if (response.status === 200) {
                    //delete cookie & auth state false
                    document.cookie = 'session=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                    document.cookie = 'user_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                    localStorage.setItem('isAuth',false)
                    window.location.replace('/')
                } else {
                    console.log('error logout')
                }
            }
    }
    async getHtml() {
        //get -> set profile data -> by id
        //fetch - post, edit -> name etc
        return super.showHeader('auth');
    }
}