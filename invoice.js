var item1, item2, inv1;

function invoices() {
    console.log('hello');
}

function ItemObj(id, price, quantity, note) {
    this.id = id;
    this.price = price;
    this.quantity = quantity;
    this.note = note;
}

item1 = new ItemObj('ochestra224', 125, 2, "Center Mezzanine");
item2 = new ItemObj('balcony32', 50, 1, "Left Mezzanine");

function InvoiceObj(id, customerId, raised, due, paid, note, item1, item2) {
    this.id = id;
    this.customerId = customerId;
    this.raised = raised;
    this.due = due;
    this.paid = paid;
    this.note = note;

    // assume only two items per invoice
    this.items = [item1, item2];
}

inv1 = new InvoiceObj(1, 111, "10/15/2014", "10/22/2014", true, "Jazz concert", item1, item2);

inv2 = new InvoiceObj(2, 222, "11/16/2015", "11/23/2015", false, "Jazz concert", item1, item2);

function MarshallInvoices(invoices) {
    var myLog = function(arg) {
        console.log(arg);
    };
    

    function write() {
        for (var i = 0; i < invoices.length; i++) {
            myLog(invoices[i]);
        }

    }

    write();
}

MarshallInvoices([inv1, inv2]);