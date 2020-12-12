package com.atlas.dis.entity;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Table(indexes = {
      @Index(name = "mobId", columnList = "continent")
})
public class GlobalDrop implements Serializable {
   private static final long serialVersionUID = 1L;

   @Id
   @GeneratedValue(strategy=GenerationType.IDENTITY)
   private Integer id;

   @Column(nullable = false)
   private Integer continent = -1;

   @Column(nullable = false)
   private Integer itemId = 0;

   @Column(nullable = false)
   private Integer minimumQuantity = 1;

   @Column(nullable = false)
   private Integer maximumQuantity = 1;

   @Column(nullable = false)
   private Integer questId = 0;

   @Column(nullable = false)
   private Integer chance = 0;

   @Column
   private String comments;

   public GlobalDrop() {
   }

   public Integer getId() {
      return id;
   }

   public void setId(Integer id) {
      this.id = id;
   }

   public Integer getContinent() {
      return continent;
   }

   public void setContinent(Integer continent) {
      this.continent = continent;
   }

   public Integer getItemId() {
      return itemId;
   }

   public void setItemId(Integer itemId) {
      this.itemId = itemId;
   }

   public Integer getMinimumQuantity() {
      return minimumQuantity;
   }

   public void setMinimumQuantity(Integer minimumQuantity) {
      this.minimumQuantity = minimumQuantity;
   }

   public Integer getMaximumQuantity() {
      return maximumQuantity;
   }

   public void setMaximumQuantity(Integer maximumQuantity) {
      this.maximumQuantity = maximumQuantity;
   }

   public Integer getQuestId() {
      return questId;
   }

   public void setQuestId(Integer questId) {
      this.questId = questId;
   }

   public Integer getChance() {
      return chance;
   }

   public void setChance(Integer chance) {
      this.chance = chance;
   }

   public String getComments() {
      return comments;
   }

   public void setComments(String comments) {
      this.comments = comments;
   }
}
