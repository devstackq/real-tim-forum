import Parent from "./Parent.js";
import { global } from './Signin.js'

export default class extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }
    setTitle(title) {
        document.title = title;
    }

    async init() {

        console.log(document.getElementById('#logout'), 'logout')

        document.querySelector('#logout').onclick = async function() {

            const uuid = localStorage.getItem('session');

            if (uuid != undefined || uuid != '') {
                let session = { uuid: '' }

                session.uuid = uuid

                let response = await fetch('http://localhost:6969/api/logout', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json;charset=utf-8'
                    },
                    body: JSON.stringify(session)
                });

                let result = await response.json();
                console.log(response)
                if (response.status === 200) {
                    // toggle('logout')
                    document.cookie = 'session=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                    document.cookie = 'user_id=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                    window.isAuth = false
                        // document.cookie = 'session=; Max-Age=-99999999;';
                        //delete Cookie
                    window.location.replace('/')
                } else {
                    this.showNotify(result, 'error')
                    console.log('error logout')
                }
            }
        }
    }
    async getHtml() {
        //get -> set profile data -> by id
        //fetch - post, edit -> name etc
        return super.showHeader('free');
    }
}