const User = require('./User');
const dbSyncConn = require('./db')


class Persistence {
    static RegisterUser(user) {
        // Try to enter a new record in the database.
        try {
            const queryResult = dbSyncConn.query(
                `INSERT INTO user VALUES( ${this.userToQuery(user)} )`)
            console.log(queryResult);
        } catch (error) {
            console.log("error => \n" + error);
            console.log("error.code => \n" + error.code);
            return { success: false, message: error.code }
        }
        return { success: true }
    }
    static retrieveUser({ email, password }) {
        let storedUser = null
        try {
            // returns an array. We know that there will only be a single element
            // because the email is unique
            console.log("Query, ", `SELECT * FROM user WHERE email='${email}' AND password='${password}'`);
            storedUser = dbSyncConn.query(
                `SELECT * FROM user WHERE email='${email}' AND password='${password}'`
            )
            console.log("storedUser =>", storedUser)
        } catch (error) {
            console.log(error)
            return {error}
        }
        if (storedUser[0] == null) {
            return { success: false, message: "The email/password combination is incorrect" }
        }
        // else, the user must exist
        return{succes: true, user : storedUser[0]}

    }

    // Return a string representing the query format of the object's values.
    // This string should be passed in the values() of an sql query.
    static userToQuery(user) {
        return `'${user.firstname}','${user.lastname}','${user.email}','${user.password}',${user.isteacher}`
    }
}

module.exports = Persistence;