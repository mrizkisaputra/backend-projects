import fetchGithubUserActivity from "./request.js";

/* get argument from cli */
let username = process.argv[2]
if (!username) {
    console.error("Please provide a github username!")
    process.exit(1)
}

/* fetch github user activity */
fetchGithubUserActivity(username)
.then(items => {
    displayUserActivity(items)
})
.catch(err => {
    console.error(err)
})

function displayUserActivity(items) {
    /* if user no have activities */
    if (items.length < 1) {
        console.log(`${username} No recent activity found!`)
        process.exit(1)
    }

    items.forEach((item, index) => {
        switch (item.type) {
            case "PushEvent": console.log(`- Pushed ${item.payload.commits.length} commits to ${item.repo.name}`); break;
            case "IssuesEvent": {
                let action = item.payload.action.charAt(0).toUpperCase()+item.payload.action.slice(1)
                console.log(`- ${action} a new issue in ${item.repo.name}`);
            } break;
            case "WatchEvent": console.log(`Starred ${item.repo.name}`); break;
        }
    })
}