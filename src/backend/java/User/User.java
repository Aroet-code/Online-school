package User;

import Notification.Subscriber;
import Notification.Subscription;

import java.time.ZonedDateTime;
import java.util.ArrayList;

public class User implements Subscriber {
    private String username;
    private int age;
    private ArrayList<String> courses;
    private ZonedDateTime registrationDate;
    private ArrayList<Subscription> subscriptions;
    private ArrayList<String> tags;

    public User(String username){
        this.username = username;
        this.registrationDate = ZonedDateTime.now();
    }

    public ArrayList<String> getCourses() {
        return courses;
    }

    public void addCourse(String course){
        courses.add(course);
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public String getUsername() {
        return username;
    }

    public void addTag(String tag){
        tags.add(tag);
    }

    public boolean hasTag(String tag){
        return tags.contains(tag);
    }

    public void setUsername(String username) {
        this.username = username;
    }

    @Override
    public void onSubscribe(Subscription subscription) {

    }

    public ArrayList<Subscription> getSubscriptions() {
        return subscriptions;
    }

    @Override
    public void onNotification() {

    }
}
