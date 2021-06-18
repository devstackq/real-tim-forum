export default class Parent {

    constructor(text, type) {
        this.text = text
        this.type = type
        this.value = ''
    }
    getPostById(id) {
        console.log(id, 'id')
    }

    async Fetch(endpoint, object) {
        let response = await fetch(`http://localhost:6969/api/${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(object)
        });
        if (response.status == 200) {
            let result = await response.json()
            return result, 200
        }else {
            return null, 400
        }
    }

//render posts & getPostById
    render(item, idx, where) {

        let wrapper = document.querySelector(where)
        let btn = document.createElement('button')
        let div = document.createElement('div')

        for (let [k, v] of Object.entries(item)) {

            if (v != "" && v != null) {
                let span = document.createElement('span')
                div.className = 'postWrapper'
                span.id = k
                span.textContent = ` ${k} : ${v}`

                // btn.onclick = this.getPostById & try DRY funcs

                btn.onclick = async function() {
                    
                    let postId = { id: 0 }
                    postId.id = item['id']
                   
                    // DRY - func 
                    let response = await fetch('http://localhost:6969/api/post/id', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json;charset=utf-8'
                        },
                        body: JSON.stringify(postId)
                    });
                    let res = await response.json(0)

                    if (response.status == 200) {
                        // let status, res  = this.Fetch('post/id', postId)
                    // if (status == 200) {
                    let parent = document.querySelector('.postContainer')
                        
                        parent.innerHTML = ""                    
                    
                        let btnLike = document.createElement('button')
                                        
                    let btnDislike = document.createElement('button')
            btnLike.textContent='like'
            btnDislike.textContent = 'dislike'

            btnLike.onclick= async ()=> {
            let vote = {type: 'like'}

    let response = await fetch('http://localhost:6969/api/vote', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(vote)
    });
    let res = await response.json(0)
console.log(res)
    if (response.status == 200) {
console.log('vote like')
    }
}
                    for(let [k,v] of Object.entries(res)) {
                            let span = document.createElement('span')
                            if(v != null && v != "") {
                                span.textContent = `${k} : ${v}`
                            }
                            parent.append(span)
                        }
                        parent.append(btnLike)
                        parent.append(btnDislike)
                        // window.location.replace(`/post/id`)
                            //redirect post componen   or here render () ?
                        }
                };
                btn.value = idx
                btn.textContent = `see post`
                div.appendChild(span)
            }
            div.appendChild(btn)
            wrapper.appendChild(div)
            }}    

    showHeader(type) {

        let login = ""
        let register = ""
        let logout = ""
        let profile = ""

        if (type == 'free') {
            profile = ""
            logout = ""
            register = `<a href="/signup"  class="nav__link signup" data-link>Signup</a>`
            login = `<a href="/signin"  class="nav__link signin" data-link>Signin</a>`
        } else if (type == 'auth') {
            register = ""
            login = ""
            logout = `<a href="/logout" id='logout' class="nav__link logout" data-link>Logout</a>`
            profile = `<a href="/profile" class="nav__link" data-link>Profile</a>`
        }

        return `
        <nav class="nav">
        <a href="/all" class="nav__link" data-link>Main</a>
        ${login}
        ${register}
        <div class="dropdown">
          <button class="dropbtn">Categories</button>
          <div class="dropdown-content">
          <a href="/love" data-link>love</a>
          <a href="/science" data-link>science</a>
          <a href="/nature" data-link>nature</a>
        </div>
        </div>
       ${profile}
        ${logout}
    </nav>
    <span class='notify' > </span> 
`
    }

    showNotify(text, type) {

        let notify = document.getElementsByClassName('notify')[0]

        if (type == 'error') {
            notify.style.display = 'block'
            notify.textContent = text
        } else if (type == 'hide') {
            notify.style.display = 'none'
        }
    }
}