const {Get} = require('./CreateTest')

Get('/','message should say healthy',(expect, res) =>{
    expect(res.body.response).to.equal('hello world')
    expect(res.body.method).to.equal('GET')
})
