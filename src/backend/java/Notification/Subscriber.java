package Notification;

import java.util.ArrayList;

public interface Subscriber {
    void onSubscribe(Subscription subscription);
    void onNotification();
    ArrayList<Subscription> getSubscriptions();
}
