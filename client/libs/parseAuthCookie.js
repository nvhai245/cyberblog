import cookie from 'cookie'
function parseAuthCookie(req) {
    return cookie.parse(req ? req.headers.cookie || "" : document.cookie)
}

export default parseAuthCookie