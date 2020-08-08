import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Doorman',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: Home(),
    );
  }
}

class Home extends StatefulWidget {
  @override
  _HomeState createState() => _HomeState();
}

class _HomeState extends State<Home> {
  // TODO: change this to configuration
  final String hostUrl = 'http://192.168.1.6:8080';

  bool _locked = true;

  void _unlock() async {
    setState(() {
      _locked = false;
    });
    try {
      await http.get(hostUrl + '/unlock').timeout(Duration(seconds: 2));
      setState(() {
        _locked = true;
      });
    } catch (exception) {
      print(exception);
      setState(() {
        _locked = true;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Center(
            child: AnimatedOpacity(
                opacity: _locked ? 1.0 : 0.0,
                duration: Duration(milliseconds: 500),
                child: IconButton(
                    icon: _locked ? Icon(Icons.lock) : Icon(Icons.lock_open),
                    iconSize: 45,
                    tooltip: 'Unlock',
                    onPressed: _locked ? _unlock : null
                )
            )
        )
    );
  }
}
