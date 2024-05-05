document.addEventListener('DOMContentLoaded', function() {
    const orderForm = document.getElementById('orderForm');
    const orderIdInput = document.getElementById('orderId');

    orderForm.addEventListener('submit', async function(event) {
        event.preventDefault();

        const orderId = orderIdInput.value;
        const url = `http://localhost:8080/api/order/${orderId}`;

        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`HTTP error status: ${response.status}`);
            }
            const data = await response.json();

            document.getElementById('orderUid').textContent = data.order_uid;
            document.getElementById('trackNumber').textContent = data.track_number;
            document.getElementById('entry').textContent = data.entry;
            document.getElementById('locale').textContent = data.locale;
            document.getElementById('internalSignature').textContent = data.internal_signature;
            document.getElementById('customerId').textContent = data.customer_id;
            document.getElementById('deliveryService').textContent = data.delivery_service;
            document.getElementById('shardkey').textContent = data.shardkey;
            document.getElementById('smId').textContent = data.sm_id;
            document.getElementById('dateCreated').textContent = data.date_created;
            document.getElementById('oofShard').textContent = data.oof_shard;

            document.getElementById('deliveryName').textContent = data.delivery.name;
            document.getElementById('deliveryPhone').textContent = data.delivery.phone;
            document.getElementById('deliveryZip').textContent = data.delivery.zip;
            document.getElementById('deliveryCity').textContent = data.delivery.city;
            document.getElementById('deliveryAddress').textContent = data.delivery.address;
            document.getElementById('deliveryRegion').textContent = data.delivery.region;
            document.getElementById('deliveryEmail').textContent = data.delivery.email;

            document.getElementById('transactionId').textContent = data.payment.transaction;
            document.getElementById('requestId').textContent = data.payment.request_id;
            document.getElementById('currency').textContent = data.payment.currency;
            document.getElementById('provider').textContent = data.payment.provider;
            document.getElementById('amount').textContent = data.payment.amount;
            document.getElementById('paymentDate').textContent = new Date(data.payment.payment_dt * 1000).toLocaleString();
            document.getElementById('bank').textContent = data.payment.bank;
            document.getElementById('deliveryCost').textContent = data.payment.delivery_cost;
            document.getElementById('goodsTotal').textContent = data.payment.goods_total;
            document.getElementById('customFee').textContent = data.payment.custom_fee;

            const itemsList = document.getElementById('itemsList');
            data.items.forEach(item => {
                const itemBlock = document.createElement('div');
                itemBlock.className = 'itemBlock';

                const chrtIdElement = document.createElement('p');
                chrtIdElement.textContent = `ChrtId: ${item.chrt_id}`;
                const trackNumberElement = document.createElement('p');
                trackNumberElement.textContent = `TrackNumber: ${item.track_number}`;
                const priceElement = document.createElement('p');
                priceElement.textContent = `Price: ${item.price}`;
                const ridElement = document.createElement('p');
                ridElement.textContent = `Rid: ${item.rid}`;
                const nameElement = document.createElement('p');
                nameElement.textContent = `Name: ${item.name}`;
                const saleElement = document.createElement('p');
                saleElement.textContent = `Sale: ${item.sale}`;
                const sizeElement = document.createElement('p');
                sizeElement.textContent = `Size: ${item.size}`;
                const totalPriceElement = document.createElement('p');
                totalPriceElement.textContent = `TotalPrice: ${item.total_price}`;
                const nmIdElement = document.createElement('p');
                nmIdElement.textContent = `NmId: ${item.nm_id}`;
                const brandElement = document.createElement('p');
                brandElement.textContent = `Brand: ${item.brand}`;
                const statusElement = document.createElement('p');
                statusElement.textContent = `Status: ${item.status}`;

                itemBlock.appendChild(chrtIdElement);
                itemBlock.appendChild(trackNumberElement);
                itemBlock.appendChild(priceElement);
                itemBlock.appendChild(ridElement);
                itemBlock.appendChild(nameElement);
                itemBlock.appendChild(saleElement);
                itemBlock.appendChild(sizeElement);
                itemBlock.appendChild(totalPriceElement);
                itemBlock.appendChild(nmIdElement);
                itemBlock.appendChild(brandElement);
                itemBlock.appendChild(statusElement);

                itemsList.appendChild(itemBlock);
            });

            console.log(data);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    });
});