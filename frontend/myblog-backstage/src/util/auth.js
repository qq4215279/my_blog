import config from "@/config/website.js"

let inFifteenMinutes = new Date(new Date().getTime() + 120 * 60 * 1000);
const TokenKey = config.accessTokenKey;
const Cookies = {
    set(name,value,expires){
        document.cookie = name+"="+value+";expires="+expires;
    },
    get(name){
        var str = document.cookie;
        var arr = str.split("; ");
        for(var i = 0; i < arr.length; i++){
            //console.log(arr[i]);
            var newArr = arr[i].split("=");
            if(newArr[0]==name){
                return newArr[1];
            }
        }
    },
    remove(name){
        this.set(name,"",-1);
    }
}
export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token, inFifteenMinutes)
}

export function removeToken() {
    return Cookies.remove(TokenKey)
}