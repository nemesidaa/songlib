package model

import (
	"time"
)

/*
	Parts:
		- {
 			"group": "Muse",
 			"song": "Supermassive Black Hole"
		}
		- openapi: 3.0.3
			info:
  			title: Music info
  			version: 0.0.1
			paths:
  			/info:
    			get:
      			parameters:
        			- name: group
          				in: query
          				required: true
          				schema:
            				type: string
        			- name: song
          				in: query
          				required: true
          				schema:
            				type: string
      			responses:
        			'200':
          				description: Ok
          				content:
            				application/json:
              				schema:
                				$ref: '#/components/schemas/SongDetail'
        			'400':
          				description: Bad request
        			'500':
          				description: Internal server error
			components:
  				schemas:
    				SongDetail:
      					required:
        					- releaseDate
        					- text
       			  			- link
      						type: object
	  					$properties:
        					releaseDate:
          						type: string
          						example: 16.07.2006
        					text:
          						type: string
          						example: Ooh baby, don't you know I suffer?
        					patronymic:
          						type: string
          						example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
*/

type Song struct {
	ID    string `gorm:"primaryKey"`
	Group string `gorm:"not null"`
	Song  string `gorm:"not null"`
	// as additional, from the external API
	ReleaseDate time.Time `gorm:"type:timestamp without time zone;not null"`
	Text        string    `gorm:"not null"`
	Link        string    `gorm:"not null"`
}

func (s *Song) TableName() string {
	return "songs"
}
