async function fetchGithubUserActivity(username) {
    let domain = "https://api.github.com"
    let response = await fetch(`${domain}/users/${username}/events`, {
        headers: {
            "Accept": "application/vnd.github+json"
        }
    })
    if (!response.ok) {
        switch (response.status) {
            case 304: throw new Error(`Not modified - ${response.status}`); break;
            case 403: throw new Error(`Forbidden - ${response.status}`); break;
            case 404: throw new Error(`Username not found - ${response.status}`); break;
            case 503: throw new Error(`Service unavailable - ${response.status}`); break;
            default: throw new Error(`Server Internal Error - ${response.status}`); break;
        }
    }
    return response.json()
}

export default fetchGithubUserActivity