import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface IAuthor {
  firstName: string,
  middleName: string,
  lastName: string
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'ui';
  public authors: IAuthor[] = [];

  constructor(private httpClient: HttpClient) {}

  // ngOnInit() {
  //   this.getAuthors();
  // }

  ngOnInit() {
    this.httpClient.get<IAuthor[]>('/api/authors')
                        .subscribe({
                          next: data => { 
                            this.authors = data;
                          },
                          error: error => {
                            console.log('There was an error!', error);
                          }
                        })
  }
}

// {
//   'firstName': 'Benny',
//   'middleName': 'Warchild',
//   'lastName': 'Engstrom'
// },
// {
//   'firstName': 'Goose',
//   'middleName': 'Warchild',
//   'lastName': 'Engstrom'
// }