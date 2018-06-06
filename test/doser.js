module.exports = {
  Create: function(n){
    n.click('li#react-tabs-14')
    .wait(500)
    .click('input#add_doser')
    .wait(500)
    .type('input#doser_name','Two Part - CaCO3')
    .wait(500)
    .click('button#new_doserjack')
    .wait(500)
    .click('span#new_doser-J0')
    .wait(500)
    .click('input#create_pump')
    .wait(500)
    .click('input#schedule-pump-1')
    .wait(500)
    .click('input#pump-enable-1')
    .wait(500)
    .type('input#day')
    .type('input#day', '*')
    .wait(500)
    .type('input#hour')
    .type('input#hour', '1,9,17')
    .wait(500)
    .type('input#minute')
    .type('input#minute', '1')
    .wait(500)
    .type('input#second')
    .type('input#second', '1')
    .wait(500)
    .type('input#set-duration-1')
    .type('input#set-duration-1', '15')
    .wait(500)
    .click('input#set-schedule-1')
    .wait(1500)

    return(function(){
      return('Doser created')
    })
  }
}

